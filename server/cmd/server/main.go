package main

import (
	"../../balancers"
	"../../db"
	"flag"
)

var httpPortNumber = flag.Int("p", 3001, "HTTP port number")

func main() {
	flag.Parse()

	connect := db.Connect{
		Host : "localhost",
		Port:  5432,
		User:  "postgres",
		Password: "1234",
		Dbname:   "lab3",
	}
	DB := connect.OpenDB()

	balancers.Composer(*httpPortNumber, DB)

}