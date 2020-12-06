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
	NotUsedMachines []int `json:"notUsedMachines"`
}

func HandleListBalancers(res http.ResponseWriter, db *sql.DB) ([]BalancerInf, error){
	queryString := fmt.Sprintf(`
	SELECT CASE WHEN a.id is NOT NULL THEN a.id ELSE b.id END, a.used, b.notUsed FROM (
		SELECT balancer_id AS "id",
       		array_agg(machine_id) AS used
		FROM ConnectToBalancers, Machines
		WHERE ConnectToBalancers.machine_id = Machines.id AND Machines.isUsed = true
		GROUP BY balancer_id
		) as a
		full join (
		SELECT balancer_id AS "id", array_agg(machine_id) AS notUsed
        		FROM ConnectToBalancers, Machines
        		WHERE ConnectToBalancers.machine_id = Machines.id AND Machines.isUsed = false
        		GROUP BY balancer_id) as b on a.id = b.id
        		ORDER BY a.id, b.id;`)

	rows, queryErr := db.Query(queryString)
	if queryErr != nil {return nil, queryErr}

	var result []BalancerInf
	for rows.Next() {
		var blc BalancerInf
		var UsedMachinesInASCII, NotUsedMachinesInASCII []uint8

		rowErr := rows.Scan(&blc.Id, &UsedMachinesInASCII, &NotUsedMachinesInASCII)
		if rowErr != nil { return nil, rowErr }

		var convertErr error
		blc.UsedMachines, convertErr = tools.ASCIItoIntArr(UsedMachinesInASCII)
		if convertErr != nil {return nil, convertErr}
		blc.NotUsedMachines, convertErr = tools.ASCIItoIntArr(NotUsedMachinesInASCII)
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
