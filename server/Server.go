package server

import (
	"os"
	"net/http"
	"log"
	"github.com/gorilla/mux"
	. "trackingApp/controllers/user"
)

type Server struct {
	Router *mux.Router
}


func (a *Server) CreateServer(){
	a.Router = mux.NewRouter().StrictSlash(true)
	a.CreateRoutes()

}

func (a *Server) StartServer() {

	port := os.Getenv("PORT")

	if port == ""{
		port = "8080"
	}
	
	log.Printf("\n Server starting on port %v",port)
	log.Fatal(http.ListenAndServe(":"+port,a.Router))
}

func (a *Server) CreateRoutes() {
	var uc UserControllerBase = UserController{}
	a.Router.HandleFunc("/stores",uc.Post)
	
	// a.Router.HandleFunc("/stores",getUserData).Methods("GET")
}