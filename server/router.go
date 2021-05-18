package server

func MapRouter(server *Server) {
	server.Server.GET("/ping", Ping)
}
