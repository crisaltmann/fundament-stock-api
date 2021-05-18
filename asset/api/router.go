package asset_api

import "github.com/crisaltmann/fundament-stock-api/server"

func MapRouter(server *server.Server, handler *Handler) {
	server.Server.GET(Path, handler.GetAllAssets)
	server.Server.POST(Path, handler.InsertAsset)
	server.Server.PUT(Path+"/:id", handler.UpdateAsset)
	server.Server.GET(Path+"/:id", handler.GetById)
}
