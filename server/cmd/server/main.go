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
		Dbname:   "lab3",
	}

	DB := connect.OpenDB()

	balancers.Composer(8795, DB)

}