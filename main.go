package main

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"go-casbin/pkg/helper/slice"
	"regexp"
)

func main() {
	e, err := casbin.NewEnforcer("./configs/authz/authz_model.conf", "./configs/authz/authz_policy.csv")
	if err != nil {
		panic(err)
	}

	r := gin.Default()

	// 注册 Casbin 中间件
	r.Use(func(c *gin.Context) {
		role := ""
		if c.Request.Method == "GET" {
			role = c.Query("role")
		}
		if role == "" {
			c.AbortWithStatusJSON(403, gin.H{"message": "无效的Role"})
			return
		}
		fmt.Printf("role:%v\n", role)
		// 假设您已经在身份验证中确定了用户的角色
		c.Set("sub", role)
		c.Next()
	})

	// 使用 Casbin 中间件进行授权检查
	r.GET("/student/:obj", CasbinMiddleware(e), func(c *gin.Context) {
		// 处理学生的路由
		c.JSON(200, gin.H{
			"message": "可以访问该学生资源",
			"sub":     c.GetString("sub"),
		})
	})

	r.GET("/teacher/:obj", CasbinMiddleware(e), func(c *gin.Context) {
		// 处理教师的路由
		c.JSON(200, gin.H{
			"message": "可以访问该教师资源",
			"sub":     c.GetString("sub"),
		})
	})

	r.GET("/admin/:obj", CasbinMiddleware(e), func(c *gin.Context) {
		// 处理管理员的路由
		c.JSON(200, gin.H{
			"message": "管理员可以访问该资源",
			"sub":     c.GetString("sub"),
		})
	})

	// 启动 Gin 服务器
	r.Run(":8080")
}

func CasbinMiddleware(e *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		sub := c.GetString("sub") // 获取用户角色, 从 JWT/Cookie/Session 中获取, 此处简化了处理

		// 正则表达式模式, 匹配 /admin/xxx, /student/xxx, /teacher/xxx作为策略的obj
		pattern := `admin|student|teacher`

		// 编译正则表达式
		re := regexp.MustCompile(pattern)

		// 查找所有匹配项
		matches := re.FindAllString(c.FullPath(), -1)

		// 打印匹配的路径段
		for _, match := range matches {
			fmt.Println(match)
		}

		obj := matches[0] // 获取请求的对象（路由）

		act := ""
		method := c.Request.Method
		methods := []string{"POST", "PUT", "PATCH"}

		if method == "GET" {
			act = "read"
		} else if slice.Include(methods, method) {
			act = "write"
		}
		fmt.Printf("sub:%v\n", sub)

		// 执行 Casbin 授权检查
		res, err := e.Enforce(sub, obj, act)
		fmt.Printf("res:%v\n", res)
		if err != nil || res == false {
			c.AbortWithStatusJSON(403, gin.H{"message": "无权访问该资源"})
			return
		}

		c.Next()
	}
}
