package balancers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"../tools"
)

type BalancerInf struct {
	Id int `json:"id"`
	UsedMachenes []int `json:"usedMachenes"`
	TotalMachenes int `json:"totalMachenes"`
}

func HandleListBalancers(res http.ResponseWriter, db *sql.DB) ([]BalancerInf, error){
	queryString := fmt.Sprintf(`
		SELECT balancer_id AS "id",
       		array_agg(machine_id) AS "usedMachenes",
      		COUNT(*) AS "totalMachenes"
		FROM ConnectToBalancers, Machines
		WHERE ConnectToBalancers.machine_id = Machines.id AND Machines.isUsed = true
		GROUP BY balancer_id;`)
	rows, queryErr := db.Query(queryString)
	if queryErr != nil {return nil, queryErr}

	var result []BalancerInf
	for rows.Next() {
		var blc BalancerInf
		var UsedMachenesInASCII []uint8
		rowErr := rows.Scan(&blc.Id, &UsedMachenesInASCII, &blc.TotalMachenes)
		if rowErr != nil {panic(rowErr)}
		var convertErr error
		blc.UsedMachenes, convertErr = tools.ASCIItoIntArr(UsedMachenesInASCII)
		if convertErr != nil {panic(convertErr)}
		result = append(result, blc)
	}
	return result, nil
}

func balancersCloser(db *sql.DB) func (res http.ResponseWriter, req *http.Request) {
	return func (res http.ResponseWriter, req *http.Request) {
		if req.Method == "GET" {
			res.Header().Set("Access-Control-Allow-Origin", "*")
			result, errBlc := HandleListBalancers(res, db)
			if errBlc != nil {panic(errBlc)}

			json, err := json.Marshal(result)

			if err != nil {
				http.Error(res, err.Error(), http.StatusInternalServerError)
				return
			}

			res.Header().Set("Content-Type", "application/json")
			res.Write(json)
		}
	}
}
