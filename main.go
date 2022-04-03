package main

import (
	"trackingApp/server"
)


func main(){
	var app server.ServerBase = &server.Server{}
	app.CreateServer()
	app.CreateRoutes()
	app.StartServer()
}