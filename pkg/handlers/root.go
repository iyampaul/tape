package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"tape/pkg/actions"
	"tape/pkg/models"

	"github.com/gorilla/mux"
)

func root(w http.ResponseWriter, r *http.Request, actions actions.ActionList) {

	for _, action := range actions {
		if mux.Vars(r)["route"] == action.Obj.Route {
			if action.Authenticate(mux.Vars(r)["key"]) || action.Obj.Key == "" {
				reqParsed, err := parse(r.Body)
				if err != nil {
					// Event error log
					w.Write(genResponse(false, reqParsed["Error"]))
					return
				} else {
					success, response := action.Execute(reqParsed)
					w.Write(genResponse(success, response))
					return
				}
			} else {
				log.Printf("Event: Authentication Failed")
				w.Write(genResponse(false, "Authentication Failed"))
				return
			}
		}
	}
}

func parse(input io.ReadCloser) (models.RequestData, error) {
	requestRaw, err := io.ReadAll(input)
	if err != nil {
		return parseError("Event: Unable to read request body", err), err
	}

	reqBody := models.RequestData{}
	err = json.Unmarshal(requestRaw, &reqBody)
	if err != nil {
		return parseError("Event: Unable to parse JSON body", err), err
	} else {
		return reqBody, err
	}

}
