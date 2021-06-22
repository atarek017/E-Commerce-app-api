package main

import (
	"database/sql"
	"fmt"
)

// SuccesResponse Struct (Model)
type SuccesResponse struct {
	Code    int32       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// FailResponse Struct (Model)
type FailResponse struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

type User struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Photo      string `json:"photo"`
	Address    string `json:"address"`
	Phone      string `json:"phone"`
	Verify     string `json:"verify"`
	IsVerifyed string `json:"isverifyed"`
}

//opens a connection to database using its name and its server username and password
func openDatabaseConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/soq")
	if err != nil {
		fmt.Println("Error opening database..")
		fmt.Println(err.Error())
		panic(err.Error())
		return nil
	}
	return db
}
