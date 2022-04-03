package server

import (
	"os"
	"net/http"
	"fmt"
	"log"
	"github.com/gorilla/mux"
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
	a.Router.HandleFunc("/stores",func(w http.ResponseWriter, req *http.Request){
		fmt.Fprint(w,"Hello There 2")
	})
	
	// a.Router.HandleFunc("/stores",getUserData).Methods("GET")
}