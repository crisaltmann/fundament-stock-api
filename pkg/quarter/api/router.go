package quarter_api

import "github.com/crisaltmann/fundament-stock-api/server"

func MapRouter(server *server.Server, handler Handler) {
	server.Server.GET(Path, handler.GetQuarters)
	server.Server.GET(Path+"/:id", handler.GetQuarter)
}
