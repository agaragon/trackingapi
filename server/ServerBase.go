package server

type ServerBase interface {
	StartServer()
	CreateServer()
	CreateRoutes()
}