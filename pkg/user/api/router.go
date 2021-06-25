package user_api

import "github.com/crisaltmann/fundament-stock-api/server"

// @title Swagger Example API
// @version 1.0
// @description This is a sample server asset server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host another.swagger.io
// @BasePath /v2
func MapRouter(server *server.Server, handler Handler) {
	server.Server.POST(Path_Login, handler.Login)
}
