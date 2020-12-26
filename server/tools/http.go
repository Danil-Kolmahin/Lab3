package tools

import (
	"encoding/json"
	"log"
	"net/http"
)

type errorObject struct {
	Message string `json:"message"`
}

// WriteJsonOk sends 400 error response with a JSON object describing the error reason.
func WriteJsonBadRequest(rw http.ResponseWriter, message string) {
	writeJson(rw, http.StatusBadRequest, &errorObject{Message: message})
}

// WriteJsonOk sends 500 error response.
func WriteJsonInternalError(rw http.ResponseWriter) {
	writeJson(rw, http.StatusBadRequest, &errorObject{Message: "internal error happened"})
}

// WriteJsonOk sends 200 response to the client serializing the input object in JSON format.
func WriteJsonOk(rw http.ResponseWriter, res interface{}) {
	writeJson(rw, http.StatusOK, res)
}

func writeJson(rw http.ResponseWriter, status int, res interface{}) {
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	rw.Res.Header().Set("Content-Type", "application/json")

	rw.WriteHeader(status)
	err := json.NewEncoder(rw).Encode(res)
	if err != nil {
		log.Printf("Error writing response: %s", err)
	}
}

/*
type HttpHandler struct {
	Res http.ResponseWriter
}

func (hh *HttpHandler) HttpErrorChecker(err error){
	if err != nil {
		http.Error(hh.Res, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (hh *HttpHandler)  HttpStandartHeader() {
	hh.Res.Header().Set("Access-Control-Allow-Origin", "*")
	hh.Res.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	hh.Res.Header().Set("Content-Type", "application/json")
}

func (hh *HttpHandler) WriteJSON(result interface{}) {
	json, err := json.Marshal(result)
	hh.HttpErrorChecker(err)

	hh.Res.Write(json)
}
*/
