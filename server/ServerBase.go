package server

type ServerBase interface {
	StartServer() interface{}
	CreateServer() interface{}
	CreateRoutes() interface{}
}