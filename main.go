package main

import (
	"trackingApp/server"
)


func main(){
	app := server.Server{}
	app.CreateServer()
	app.CreateRoutes()
	app.StartServer()
}