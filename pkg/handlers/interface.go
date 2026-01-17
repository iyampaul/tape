package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"tape/pkg/actions"
	"tape/pkg/events"
	"tape/pkg/models"

	"github.com/gorilla/mux"
)

func Initialize(config *models.ConfigObj, actions actions.ActionList) {

	route := mux.NewRouter()
	route.HandleFunc("/{route}/{key}", func(w http.ResponseWriter, r *http.Request) { root(w, r, actions) })

	serveRoute(config, route)
}

func serveRoute(config *models.ConfigObj, route *mux.Router) {

	servePort := fmt.Sprintf(":%d", config.ListenPort)

	if config.TlsEnabled {
		log.Printf("Listening for HTTPS/TLS on port %v", config.ListenPort)
		err := http.ListenAndServeTLS(servePort, config.TlsCertificate, config.TlsKey, route)
		events.CheckError(err)
	} else {
		log.Printf("Listening for HTTP server on port %v", config.ListenPort)
		err := http.ListenAndServe(servePort, route)
		events.CheckError(err)
	}
}

// Standardize response formatting
func genResponse(state bool, message string) []byte {
	response, err := json.Marshal(models.Response{
		Success:  state,
		Response: message})
	if err != nil {
		// Log the error and return a generic error message
		log.Printf("Error marshaling response: %v", err)
		return []byte(`{"success":false,"response":"Internal Server Error"}`)
	}
	return response
}

// Inbound request error parsing
func parseError(msg string, err error) models.RequestData {
	log.Printf("Error: %s (%v)", msg, err.Error())
	return models.RequestData{"Error": msg}
}
