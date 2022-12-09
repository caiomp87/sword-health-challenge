package middlewares

var interceptedRoutes = map[string]string{
	"POST/v1/login/":         "public",
	"POST/v1/user/":          "create-user",
	"POST/v1/task/":          "create-task",
	"GET/v1/task/":           "list-task",
	"GET/v1/task/:id":        "get-task",
	"PATCH/v1/task/:id":      "update-task",
	"PATCH/v1/task/done/:id": "done-task",
	"DELETE/v1/task/:id":     "delete-task",
}

var technicianPermissions = []string{
	"GET/v1/task/:id",
	"GET/v1/task/",
	"POST/v1/task/",
	"PATCH/v1/task/:id",
	"PATCH/v1/task/done/:id",
}

var managerPermissions = []string{
	"POST/v1/user/",
	"GET/v1/task/",
	"GET/v1/task/:id",
	"DELETE/v1/task/:id",
}
