package main

import (
	"../../db"
	"../../balancers"
)



func main() {

	connect := db.Connect{
		Host : "localhost",
		Port:  5432,
		User:  "postgres",
		Password: "1234",
		Dbname:   "gradebook",
	}

	DB := connect.OpenDB()

	balancers.Composer(8795, DB)

	//rows, queryErr := DB.Query("SELECT * FROM balancers")
	//
	//if queryErr != nil { panic(queryErr)}

	//for rows.Next() {
	//	var id string
	//	rowErr := rows.Scan(&id)
	//	if rowErr != nil {panic(rowErr)}
	//	result = append(result, id)
	//}
	//fmt.Println(result)
	//defer db.Close()

}