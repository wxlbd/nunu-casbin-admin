package utils

import "strings"

// MenuNameToAPI 将菜单名称转换为 API 路径和方法
func MenuNameToAPI(name string) (path, method string) {
	parts := strings.Split(name, ":")
	if len(parts) < 3 {
		return "", ""
	}

	// 处理路径
	path = "/admin"
	for i := 0; i < len(parts)-1; i++ {
		path += "/" + parts[i]
	}

	// 处理方法
	action := parts[len(parts)-1]
	switch action {
	case "save":
		method = "POST"
	case "update":
		method = "PUT"
	case "delete":
		method = "DELETE"
	case "index", "list":
		method = "GET"
	default:
		method = "POST"
	}

	return path, method
}
