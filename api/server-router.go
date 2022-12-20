package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	handlers "github.com/jmr-repo/go-rest-api/handlers"
	resources "github.com/jmr-repo/go-rest-api/resources"
)

var data = ServerData{
	apiVersion: "v1",
	port:       "8080",
}

func ServerRouter() {

	// New Router Instance
	router := mux.NewRouter().StrictSlash(true)

	// Set Headers to accept only JSON requests
	router.Headers("Content-Type", "application/json")

	// Handle Error 404
	router.NotFoundHandler = http.HandlerFunc(handlers.HandlerNotFound)

	// subrouter so it can be used a version previously to any resource
	path := router.PathPrefix(data.apiVersion).Subrouter()

	// request handler resource
	path.Use(handlers.HandlerRequestHandler)

	// index resource
	path.HandleFunc("/", resources.ResourceIndex)

	// print text to let knoe the server is running
	log.Println("Listenting on Port: " + data.port)

	// start server or log error
	err := http.ListenAndServe(":"+data.port, router)

	if err != nil {
		log.Fatal("Server Start Error: " + err.Error())
	}

}
