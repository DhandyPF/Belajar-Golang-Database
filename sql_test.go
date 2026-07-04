package belajar_golang_database

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"
)

func TestExecutionSQL(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	
	script := "INSERT INTO customer(id, name, email, balance, rating, birth_date, married) VALUES(?, ?, ?, ?, ?, ?, ?)"
	_, err := db.ExecContext(ctx, script, "2", "Dhandy", "dhandy@gmail.com", 1000, 5.0, time.Now(), false)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new Customer")
}

func TestQuerySQL(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	
	script := "SELECT id, name FROM customer"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var id string
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println("ID : ", id)
		fmt.Println("Name : ", name)
	}

	defer rows.Close()
}

func TestQuerySQLComplex(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	
	script := "SELECT id, name, email, balance, rating, birth_date, married, created_at FROM customer"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var id, name string
		var email sql.NullString
		var balance sql.NullInt32
		var rating sql.NullFloat64
		var birth_date sql.NullTime
		var created_at time.Time
		var married bool
		err := rows.Scan(&id, &name, &email, &balance, &rating, &birth_date, &married, &created_at)
		if err != nil {
			panic(err)
		}

		fmt.Println("===========================================")
		fmt.Println("ID : ", id)
		fmt.Println("Name : ", name)
		if email.Valid {
			fmt.Println("Email : ", email.String)
		}
		if balance.Valid {
			fmt.Println("Balance : ", balance.Int32)
		}
		if rating.Valid {
			fmt.Println("Rating : ", rating.Float64)
		}
		if birth_date.Valid {
			fmt.Println("Birth Date : ", birth_date.Time)
		}
		fmt.Println("Married : ", married)
		fmt.Println("Created At : ", created_at)
		fmt.Println("===========================================")
	}

	defer rows.Close()
}

func TestSQLInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	
	username := "admin'; #"
	password := "yes"

	script := "SELECT username FROM user WHERE username = '" + username + "' AND password = '" + password + "' LIMIT 1"
	// fmt.Println(script)
	rows, err := db.QueryContext(ctx, script)
	if err != nil{
		panic(err)
	}

	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Berhasil Login", username)
	} else {
		fmt.Println("Gagal Login")
	}
}

func TestSQLInjectionSafe(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	
	username := "admin"
	password := "admin"

	script := "SELECT username FROM user WHERE username = ? AND password = ? LIMIT 1"
	// fmt.Println(script)
	rows, err := db.QueryContext(ctx, script, username, password)
	if err != nil{
		panic(err)
	}

	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Berhasil Login", username)
	} else {
		fmt.Println("Gagal Login")
	}
}