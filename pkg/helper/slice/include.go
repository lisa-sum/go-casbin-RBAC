package slice

func Include(array []string, item string) bool {
	for _, v := range array {
		if v == item {
			return true
		}
	}
	return false
}
