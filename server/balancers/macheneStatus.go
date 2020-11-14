package balancers

import (
	"database/sql"
	"net/http"
	"strconv"
	"fmt"
)

type MacheneStatus struct {
	MacheneId string
	IsWork bool
}

func (ms *MacheneStatus) ChangeStatus(db *sql.DB) (status string, err error){
	queryString := fmt.Sprintf("UPDATE Machines SET isUsed = %t WHERE id = %s", ms.IsWork, ms.MacheneId) //!!!!!!!!!SQL
	_, queryErr := db.Query(queryString)
	if queryErr != nil {return "error", queryErr}
	return "ok", nil
}

func statusCloser(db *sql.DB) func (res http.ResponseWriter, req *http.Request) {
	return func (res http.ResponseWriter, req *http.Request) {
		if req.Method == "POST" {
			req.ParseForm()
			isWorkStr := req.PostForm.Get("isWork")
			isWorkBool, err := strconv.ParseBool(isWorkStr)
			if err != nil {
				panic(err)
			}

			ms := MacheneStatus{
				MacheneId : req.PostForm.Get("macheneId"),
				IsWork : isWorkBool,
			}

			result, dbError := ms.ChangeStatus(db)

			if dbError != nil {panic(err)}

			res.Write([]byte(result))
		}
	}
}
