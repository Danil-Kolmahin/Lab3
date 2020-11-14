package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Connect struct {
	Host  string
	Port  int
	User  string
	Password string
	Dbname   string
}

func (c *Connect)  ConnectStr() string{
	connectStr := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		c.Host, c.Port, c.User, c.Password, c.Dbname)
	return connectStr
}

func (c *Connect) OpenDB() *sql.DB{
	db, err := sql.Open("postgres", c.ConnectStr())
	if err != nil {
		panic(err)
	}
	return db
}
