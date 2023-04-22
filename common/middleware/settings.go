package middleware

type UrlInfo struct {
	Url    string
	Method string
}

// CasbinExclude casbin 排除的路由列表

var CasbinExclude = []UrlInfo{
	{Url: "/api/login", Method: "POST"},
	{Url: "/api/logout", Method: "POST"},
	{Url: "/api/deptTree", Method: "GET"},
	{Url: "/api/menuTreeselect", Method: "GET"},
	{Url: "/api/menurole", Method: "GET"},
	{Url: "/api/menuids", Method: "GET"},
	{Url: "/api/roleMenuTreeselect/:roleId", Method: "GET"},
	{Url: "/api/roleDeptTreeselect/:roleId", Method: "GET"},
	{Url: "/", Method: "GET"},
	{Url: "/api/user/pwd/reset", Method: "PUT"},
}
