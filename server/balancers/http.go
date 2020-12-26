package balancers

import (
	"encoding/json"
	"github.com/KolmaginDanil/Lab3/server/tools"
	"log"
	"net/http"
)

// Channels HTTP handler.
type HttpHandlerFunc http.HandlerFunc

// HttpHandler creates a new instance of balancers HTTP handler.
func HttpHandler(store *Store) HttpHandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			handleListBalancers(store, rw)
		} else if r.Method == "POST" {
			handleChangeBalancerStatus(r, rw, store)
		} else {
			rw.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func handleListBalancers(store *Store, rw http.ResponseWriter) {
	res, err := store.ListBalancers()
	if err != nil {
		log.Printf("Error making query to the db: %s", err)
		tools.WriteJsonInternalError(rw)
		return
	}
	tools.WriteJsonOk(rw, res)
}

func handleChangeBalancerStatus(r *http.Request, rw http.ResponseWriter, store *Store) {
	var ms MachineStatus
	if err := json.NewDecoder(r.Body).Decode(&ms); err != nil {
		log.Printf("Error decoding balancer input: %s", err)
		tools.WriteJsonBadRequest(rw, "bad JSON payload")
		return
	}
	err := store.ChangeStatus(ms)
	if err == nil {
		tools.WriteJsonOk(rw, &ms)
	} else {
		log.Printf("Error updating : %s", err)
		tools.WriteJsonInternalError(rw)
	}
}
