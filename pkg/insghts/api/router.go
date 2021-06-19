package insight_api

import "github.com/crisaltmann/fundament-stock-api/server"

func MapRouter(server *server.Server, handler Handler) {
	server.Server.GET(Path, handler.GetInsights)
	server.Server.GET(Path+"-summary", handler.GetInsightsSummary)
}
