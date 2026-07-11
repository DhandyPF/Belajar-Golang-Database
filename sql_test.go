package belajar_golang_database

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
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

func TestExecutionSQLParameter(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "dimas'; DROP TABLE user; #"
	password := "dimas"
	
	script := "INSERT INTO user(username, password) VALUES(?, ?)"
	_, err := db.ExecContext(ctx, script, username, password)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success Insert New User")
}

func TestAutoIncrement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	email := "dimas@gmail.com"
	comment := "Dimas Ngulang Strukdat Semester 4"
	
	script := "INSERT INTO comments(email, comment) VALUES(?, ?)"
	result, err := db.ExecContext(ctx, script, email, comment)
	if err != nil {
		panic(err)
	}
	insertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	fmt.Println("Success Insert New Comment with ID", insertId)
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
	
	username := "dimas"
	password := "dimas"

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

func TestPrepareStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	script := "INSERT INTO comments(email, comment) VALUES(?, ?)"
	statement, err := db.PrepareContext(ctx, script)
	if err != nil {
		panic(err)
	}

	defer statement.Close()

	for i := 0; i < 10; i++ {
		email := "Dhandy" + strconv.Itoa(i) + "@gmail.com"
		comment := "Komentar ke " + strconv.Itoa(i)

		result, err := statement.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}

		lastInsertId, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}
		
		fmt.Println("Komen Id : ", lastInsertId)
	}
}

func TestTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	script := "INSERT INTO comments(email, comment) VALUES(?, ?)"
	// do Transaction
	for i := 0; i < 10; i++ {
		email := "Dhandy" + strconv.Itoa(i) + "@gmail.com"
		comment := "Komentar ke " + strconv.Itoa(i)

		result, err := tx.ExecContext(ctx, script, email, comment)
		if err != nil {
			panic(err)
		}

		lastInsertId, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}
		
		fmt.Println("Komen Id : ", lastInsertId)
	}

	err = tx.Commit()
	if err != nil {
		panic(err)
	}
}