package balancers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"fmt"
)

type BalancerInf struct {
	Id int `json:"id"`
	UsedMachenes []int `json:"usedMachenes"`
	TotalMachenes int `json:"totalMachenes"`
}

func HandleListBalancers(res http.ResponseWriter, db *sql.DB) ([]BalancerInf, error){
	queryString := fmt.Sprintf("%t - bool %s - string") //!!!!!!!!!SQL
	rows, queryErr := db.Query(queryString)
	if queryErr != nil {return nil, queryErr}

	var result []BalancerInf
	for rows.Next() {
		var blc BalancerInf
		rowErr := rows.Scan(&blc)
		if rowErr != nil {panic(rowErr)}
		result = append(result, blc)
	}
	//balancer1 := BalancerInf{Id:7, UsedMachenes: []int{1,2,3}, TotalMachenes: 5} //inf from DB
	//balancer2 := BalancerInf{Id:8, UsedMachenes: []int{6,7,8,9}, TotalMachenes: 15} //inf from DB
	//result := []BalancerInf{balancer1, balancer2} //DB func append to this slice
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
