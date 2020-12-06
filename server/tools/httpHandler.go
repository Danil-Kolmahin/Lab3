package tools

import (
	"encoding/json"
	"net/http"
)

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
