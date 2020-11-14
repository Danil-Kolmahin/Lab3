package balancers

import (
	"database/sql"
	"fmt"
	"net/http"
)

func Composer(port int, db *sql.DB) {
	host := fmt.Sprintf(":%d", port)
	http.HandleFunc("/getbalancers", balancersCloser(db))
	http.HandleFunc("/status", statusCloser(db))

	http.ListenAndServe(host, nil)
}