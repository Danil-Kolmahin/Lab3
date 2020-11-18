package balancers

import (
	"database/sql"
	"net/http"
	"strconv"
	"fmt"
	"../tools"
)

type MachineStatus struct {
	MachineId string
	IsWork bool
}

func (ms *MachineStatus) ChangeStatus(db *sql.DB) (status string, err error){
	queryString := fmt.Sprintf("UPDATE Machines SET isUsed = %t WHERE id = %s", ms.IsWork, ms.MachineId)
	_, queryErr := db.Query(queryString)
	if queryErr != nil {return "error", queryErr}
	return "ok", nil
}

func statusCloser(db *sql.DB) func (res http.ResponseWriter, req *http.Request) {
	return func (res http.ResponseWriter, req *http.Request) {
		if req.Method == "POST" {
			hh := tools.HttpHandler{Res: res}

			req.ParseForm()

			isWorkStr := req.PostForm.Get("isWork")
			isWorkBool, errParse := strconv.ParseBool(isWorkStr)
			hh.HttpErrorChecker(errParse)

			ms := MachineStatus{
				MachineId : req.PostForm.Get("machineId"),
				IsWork : isWorkBool,
			}
            fmt.Println(ms)
			result, dbError := ms.ChangeStatus(db)
			hh.HttpErrorChecker(dbError)

			res.Write([]byte(result))
		}
	}
}
