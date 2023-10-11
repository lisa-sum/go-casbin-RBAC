# go-casbin-RBAC

## 鉴权模型
Casbin有很多种模型, 例如:
- ACL
- RBAC
- Restful
- ABAC

本库是以RBAC模型为例, 通过Casbin的API来实现RBAC模型的鉴权功能.

### RBAC模型
一种基于角色的访问控制模型, 该模型定义了以下模型和3种角色:
subject 实体, 本库的实体为学生
object 对象, 本库的对象为学生
action 行为, 本库的行为为读和写权限

- 学生, 只能对学生路由进行读写, 其他路由无权限
- 教师, 只能对教师路由进行读写, 其他路由无权限
- 管理员, 对所有的路由进行读写

## 使用
1. 拉取本库
    ```shell
    go get github.com/lisa-sum/go-casbin-RBAC
    ```
2. 安装依赖
    ```shell
    go mod tidy
    ```

3. 运行
    ```shell
    go run main.go
    ```

4. 访问路由, 使用不同的角色访问不同的路由, 例如role=student, role=teacher, role=admin, 观察返回的结果
    1. 学生路由:
        ```shell
        curl http://localhost:8080/student/zhangsan?role=student
        ```
   2. 教师路由:
        ```shell
        curl http://localhost:8080/teacher/zhangsan?role=teacher
        ```
   3. 管理员路由:
       ```shell
       curl http://localhost:8080/admin/zhangsan?role=admin
       ```
