package functions

import (
	_ "github.com/mattn/go-sqlite3"
)

// func GetDB() (db *sql.DB) {
// 	db, _ = sql.Open("sqlite3", "user_data-database") // Open the created SQLite File
// 	defer db.Close()
// 	var tablename = "user_auth"

// 	//to get from frontend
// 	var username = "username2"
// 	var password = "password2"

// 	// Create DB
// 	// fmt.Println("Create db")
// 	// O.CreateDB()

// 	// Create Table
// 	// fmt.Println("Create table")
// 	// O.CreateAuthTable(sqliteDatabase, tablename) // Create Database Tables

// 	// INSERT RECORDS
// 	fmt.Println("insert")
// 	//hashing pass
// 	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
// 	if err != nil {
// 		// TODO: Properly handle error
// 		log.Fatal(err)
// 	}
// 	fmt.Println("Hash to store:", string(hash))
// 	error := DM.InsertAuthTable(sqliteDatabase, tablename, username, hash)
// 	if error != "None" {
// 		fmt.Println("Duplicate username")
// 	}

// 	// Authenticate user
// 	newPass := "password"
// 	auth := DM.GetPassForUser(sqliteDatabase, tablename, username, newPass)
// 	if auth == 0 {
// 		fmt.Print("User not authenticated")
// 	} else {
// 		fmt.Println("User Authenticated")
// 	}

// 	// DISPLAY INSERTED RECORDS
// 	fmt.Println("display")
// 	DM.DisplayAuthTable(sqliteDatabase, tablename)
// }
