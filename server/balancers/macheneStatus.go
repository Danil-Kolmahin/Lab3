package balancers

import (
	"github.com/KolmaginDanil/Lab3/server/tools"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

type MachineStatus struct {
	MachineId int
	IsWork bool
}

func (ms *MachineStatus) ChangeStatus(db *sql.DB) (status string, err error){
	queryString := fmt.Sprintf("UPDATE Machines SET isUsed = %t WHERE id = %d", ms.IsWork, ms.MachineId)
	_, queryErr := db.Query(queryString)
	if queryErr != nil {return "error", queryErr}
	return "ok", nil
}

func statusCloser(db *sql.DB) func (res http.ResponseWriter, req *http.Request) {
	return func (res http.ResponseWriter, req *http.Request) {
		if req.Method == "POST" {
			hh := tools.HttpHandler{Res: res}
			hh.HttpStandartHeader()

			var ms MachineStatus
			err := json.NewDecoder(req.Body).Decode(&ms)

			hh.HttpErrorChecker(err)

            fmt.Println("POST data :", ms)
			result, dbError := ms.ChangeStatus(db)
			hh.HttpErrorChecker(dbError)

			res.Write([]byte(result))
		}
	}
}
