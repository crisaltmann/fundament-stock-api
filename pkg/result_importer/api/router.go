package result_importer_api

import "github.com/crisaltmann/fundament-stock-api/server"

func MapRouter(server *server.Server, handler Handler) {
	server.Server.POST(Path, handler.Import)
	server.Server.POST(Path+"-all", handler.ImportAll)
}
