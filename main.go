package main

import (
	_ "encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage).Methods("GET")
	myRouter.HandleFunc("/soq/user/signup", signup).Methods("POST")
	myRouter.HandleFunc("/soq/user/verify", verify).Methods("POST")
	myRouter.HandleFunc("/soq/user/resendCode", resendCode).Methods("POST")
	myRouter.HandleFunc("/soq/user/login", login).Methods("POST")
	myRouter.HandleFunc("/soq/user/updateUserInfo", updateUserInfo).Methods("POST")
	myRouter.HandleFunc("/soq/products/{type}", login).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", myRouter))
}

func main() {

	handleRequests()

}
