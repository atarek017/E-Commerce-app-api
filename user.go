package main

import (
	"encoding/json"
	"github.com/segmentio/ksuid"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

func signup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	var codSend bool
	var user User
	var succesResponse SuccesResponse
	var failResponse FailResponse
	var db = openDatabaseConnection()
	_ = json.NewDecoder(r.Body).Decode(&user)

	user.ID = ksuid.New().String()
	user.Verify = strconv.Itoa(rand.Intn(9999))
	user.IsVerifyed = "0"
	//user.IsVerifyed="0"
	codSend = sendCodeSMS(user.Phone, user.Verify)

	if (codSend == false) {
		failResponse.Code = 400
		failResponse.Message = "can't send code sms pleas check your phone number"
		json.NewEncoder(w).Encode(failResponse)
		return
	}

	insert, err := db.Query("INSERT INTO soq.user VALUES (?,?,?,?,?,?,?,?,?)",
		user.ID,
		user.Name,
		user.Email,
		user.Password,
		user.Photo,
		user.Address,
		user.Phone,
		user.Verify,
		user.IsVerifyed,
	)
	defer insert.Close()
	if err != nil {
		println(err.Error())
		panic(err.Error())
		failResponse.Code = 400
		failResponse.Message = "No Data Found"
		json.NewEncoder(w).Encode(failResponse)
		return
	}
	succesResponse.Code = 200
	succesResponse.Message = "Added Successfully"
	succesResponse.Data = user
	json.NewEncoder(w).Encode(succesResponse)

}
func verify(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	var user User
	var code string
	var succesResponse SuccesResponse
	var failResponse FailResponse
	var db = openDatabaseConnection()
	_ = json.NewDecoder(r.Body).Decode(&user)

	result, err := db.Query("SELECT verify FROM soq.user WHERE id=?", user.ID, )
	defer result.Close()
	if err != nil {
		println(err.Error())
		panic(err.Error())
		failResponse.Code = 400
		failResponse.Message = "No Data Found"
		json.NewEncoder(w).Encode(failResponse)
		return
	}

	for result.Next() {
		err = result.Scan(&code)
		if err != nil {
			println(err.Error())
			panic(err.Error())
			failResponse.Code = 400
			failResponse.Message = "No Data Found"
			json.NewEncoder(w).Encode(failResponse)
			return
		}
	}

	if (code == user.Verify) {
		if setUserIsVerified(user.ID) == true {
			succesResponse.Code = 200
			succesResponse.Message = "successful verify"
			succesResponse.Data = ""
			json.NewEncoder(w).Encode(succesResponse)
		} else {
			failResponse.Code = 400
			failResponse.Message = "Failed to verify your email "
			json.NewEncoder(w).Encode(failResponse)
		}

	} else {
		failResponse.Code = 400
		failResponse.Message = "Failed check your code"
		json.NewEncoder(w).Encode(failResponse)
	}

}
func login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	var user User
	var selecteduser User
	var succesResponse SuccesResponse
	var failResponse FailResponse
	var db = openDatabaseConnection()
	_ = json.NewDecoder(r.Body).Decode(&user)

	results, err := db.Query("SELECT * FROM soq.user WHERE email = ? AND password = ? And isverifyed='1'", user.Email, user.Password)
	if err != nil {
		println("Errro   : " + err.Error())
		failResponse.Code = 400
		failResponse.Message = "check your email and password"
		json.NewEncoder(w).Encode(failResponse)
		return
	}

	for results.Next() {
		err = results.Scan(&selecteduser.ID, &selecteduser.Name,
			&selecteduser.Email, &selecteduser.Password, &selecteduser.Photo,
			&selecteduser.Address, &selecteduser.Phone, &selecteduser.Verify, &selecteduser.IsVerifyed)
		if err != nil {
			failResponse.Code = 400
			failResponse.Message = "Cannot find this user, check your email and password and try again."
			json.NewEncoder(w).Encode(failResponse)
			return
		}

	}
	if (user.Email == selecteduser.Email && user.Password == selecteduser.Password && selecteduser.Email != "") {
		succesResponse.Code = 200
		succesResponse.Message = "sign in Successfully"
		succesResponse.Data = selecteduser
		json.NewEncoder(w).Encode(succesResponse)
	} else {
		failResponse.Code = 400
		failResponse.Message = "Cannot find this user, check your email and password and try again."
		json.NewEncoder(w).Encode(failResponse)
	}

}
func resendCode(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	var user User
	var succesResponse SuccesResponse
	var failResponse FailResponse
	var codSend bool
	var db = openDatabaseConnection()
	_ = json.NewDecoder(r.Body).Decode(&user)

	user.Verify = strconv.Itoa(rand.Intn(9999))
	codSend = sendCodeSMS(user.Phone, user.Verify)
	if (codSend == false) {
		failResponse.Code = 400
		failResponse.Message = "can't send code sms pleas check your phone number"
		json.NewEncoder(w).Encode(failResponse)
		return
	}

	result, err := db.Exec("UPDATE `soq`.`user` SET `verify`=?  WHERE `id`=?", user.Verify, user.ID, )
	if err != nil {
		failResponse.Code = 400
		failResponse.Message = "can't send code sms pleas check your phone number"
		json.NewEncoder(w).Encode(failResponse)

		return
	}

	if err != nil {
		log.Fatal(err)
		failResponse.Code = 400
		failResponse.Message = "can't send code sms pleas check your phone number"
		json.NewEncoder(w).Encode(failResponse)
		return

	}
	rows, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
		failResponse.Code = 400
		failResponse.Message = "can't send code sms pleas check your phone number"
		json.NewEncoder(w).Encode(failResponse)
		return
	}
	if rows != 1 {
		log.Fatalf("expected to affect 1 row, affected %d", rows)
		failResponse.Code = 400
		failResponse.Message = "can't send code sms pleas check your phone number"
		json.NewEncoder(w).Encode(failResponse)
		return
	}

	succesResponse.Code = 200
	succesResponse.Message = "successful code sent"
	succesResponse.Data = user
	json.NewEncoder(w).Encode(succesResponse)

}
func setUserIsVerified(id string) bool {
	var db = openDatabaseConnection()

	result, err := db.Exec("UPDATE `soq`.`user` SET `isverifyed`='1'  WHERE `id`=?", id)
	if err != nil {
		return false;
	}

	if err != nil {
		log.Fatal(err)
		return false

	}
	rows, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
		return false
	}
	if rows != 1 {
		log.Fatalf("expected to affect 1 row, affected %d", rows)

		return false
	}
	return true
}
func updateUserInfo(w http.ResponseWriter, r *http.Request) {
	var user User
	var succesResponse SuccesResponse
	var failResponse FailResponse
	var db = openDatabaseConnection()
	_ = json.NewDecoder(r.Body).Decode(&user)
	result, err := db.Exec("UPDATE `soq`.`user` SET `name`=?, `email`=?, `password`=?, `photo`=?, `address`=?,`phone`=? WHERE `id`=?",
		user.Name, user.Email,
		user.Password, user.Photo,
		user.Address, user.Phone,
		user.ID)
	if err != nil {
		failResponse.Code = 400
		failResponse.Message = "can't update your profile  info"
		json.NewEncoder(w).Encode(failResponse)
		return;
	}

	if err != nil {
		log.Fatal(err)
		failResponse.Code = 400
		failResponse.Message = "can't update your profile  info"
		json.NewEncoder(w).Encode(failResponse)
		return

	}
	rows, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
		failResponse.Code = 400
		failResponse.Message = "can't update your profile  info"
		json.NewEncoder(w).Encode(failResponse)
		return
	}
	if rows != 1 {
		log.Fatalf("expected to affect 1 row, affected %d", rows)
		failResponse.Code = 400
		failResponse.Message = "can't update your profile  info"
		json.NewEncoder(w).Encode(failResponse)
		return
	}

	user=getUserData(user.ID)
	succesResponse.Code = 200
	succesResponse.Message = "Updated Successfully"
	succesResponse.Data = user
	json.NewEncoder(w).Encode(succesResponse)

}
func sendCodeSMS(phone string, code string) bool {
	resp, err := http.Get("http://world.msg91.com/api/sendhttp.php?authkey=4552ASrJPOrK5db232f0&mobiles=" + phone + "&message=your verify code is " + code + "&sender=ABCDEF&route=4&country=0")

	if (err != nil) {
		print("Eroor")

		return false
	}
	print(resp)

	return true
}

func getUserData(id string) User {
	var selecteduser User
	var db = openDatabaseConnection()

	results, err := db.Query("SELECT * FROM soq.user WHERE id=?", id)
	if err != nil {
		return User{}
	}

	for results.Next() {
		err = results.Scan(&selecteduser.ID, &selecteduser.Name,
			&selecteduser.Email, &selecteduser.Password, &selecteduser.Photo,
			&selecteduser.Address, &selecteduser.Phone, &selecteduser.Verify, &selecteduser.IsVerifyed)
		if err != nil {
			return User{}
		}
	}

	return selecteduser
}
