package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"

)

const(
	host = "localhost"
	port = 3306
	user = "postgres"
	password =" "
	dbname = "dbejemplo"
)

func main(){

	psqlInfo := fmt.Sprintf("host = %s port = %d user =%s" + "password = %s dbname = %s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil{
		panic(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil{
		panic(err)
	}


}