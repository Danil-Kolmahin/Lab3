package balancers

import (
	"../tools"
	"database/sql"
	"fmt"
	"net/http"
)

type BalancerInf struct {
	Id int `json:"id"`
	UsedMachines []int `json:"usedMachines"`
	TotalMachines int `json:"totalMachines"`
}

func HandleListBalancers(res http.ResponseWriter, db *sql.DB) ([]BalancerInf, error){
	queryString := fmt.Sprintf(`
		SELECT balancer_id AS "id",
       		array_agg(machine_id) AS "usedMachines",
      		COUNT(*) AS "totalMachines"
		FROM ConnectToBalancers, Machines
		WHERE ConnectToBalancers.machine_id = Machines.id AND Machines.isUsed = true
		GROUP BY balancer_id;`)

	rows, queryErr := db.Query(queryString)
	if queryErr != nil {return nil, queryErr}

	var result []BalancerInf
	for rows.Next() {
		var blc BalancerInf
		var UsedMachinesInASCII []uint8

		rowErr := rows.Scan(&blc.Id, &UsedMachinesInASCII, &blc.TotalMachines)
		if rowErr != nil { return nil, rowErr }

		var convertErr error
		blc.UsedMachines, convertErr = tools.ASCIItoIntArr(UsedMachinesInASCII)
		if convertErr != nil {return nil, convertErr}

		result = append(result, blc)
	}
	return result, nil
}

func balancersCloser(db *sql.DB) func (res http.ResponseWriter, req *http.Request) {
	return func (res http.ResponseWriter, req *http.Request) {
		if req.Method == "GET" {
			hh := tools.HttpHandler{Res : res}
			hh.HttpStandartHeader()

			result, errBlc := HandleListBalancers(res, db)
			hh.HttpErrorChecker(errBlc)

			hh.WriteJSON(result)
		}
	}
}
