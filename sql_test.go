package belajar_golang_database

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestExecutionSQL(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	
	script := "INSERT INTO customer(id, name, email, balance, rating, birth_date, married) VALUES('1', 'Dhandy', 'dhandy@gmail.com', 1000, 5.0, '2007-03-10', FALSE)"
	_, err := db.ExecContext(ctx, script)
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
		var id, name, email string
		var balance int32
		var rating float64
		var birth_date, created_at time.Time
		var married bool
		err := rows.Scan(&id, &name, &email, &balance, &rating, &birth_date, &married, &created_at)
		if err != nil {
			panic(err)
		}

		fmt.Println("===========================================")
		fmt.Println("ID : ", id)
		fmt.Println("Name : ", name)
		fmt.Println("Email : ", email)
		fmt.Println("Balance : ", balance)
		fmt.Println("Rating : ", rating)
		fmt.Println("Birth Date : ", birth_date)
		fmt.Println("Married : ", married)
		fmt.Println("Created At : ", created_at)
		fmt.Println("===========================================")
	}

	defer rows.Close()
}