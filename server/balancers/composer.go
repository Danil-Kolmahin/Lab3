package balancers

import (
	"database/sql"
	"fmt"
	"net/http"
)

func Composer(port int, db *sql.DB) {
	apiStatus := statusCloser(db)
	apiBalancers := balancersCloser(db)
	host := fmt.Sprintf(":%d", port)
	http.HandleFunc("/getbalancers", apiBalancers)
	http.HandleFunc("/status", apiStatus)

	http.ListenAndServe(host, nil)
}