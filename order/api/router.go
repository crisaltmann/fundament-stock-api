package order_api

import "github.com/crisaltmann/fundament-stock-api/server"

func MapRouter(server *server.Server, handler *Handler) {
	server.Server.POST(Path, handler.InsertOrder)
}
