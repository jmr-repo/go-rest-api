package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jmr-repo/go-rest-api/interfaces"
)

func HandlerNotFound(w http.ResponseWriter, r *http.Request) {

	httpStatus := http.StatusNotFound

	response = &interfaces.IDefaultResponse{
		Status:  httpStatus,
		Message: "Error 404, your request was not found",
	}

	jsonResponse, err := json.Marshal(response)

	if err != nil {
		log.Fatal("Error 404: " + err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)
	w.Write(jsonResponse)

}
