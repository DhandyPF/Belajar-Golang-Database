package belajar_golang_database

import (
	"context"
	"fmt"
	"testing"
)

func TestExecutionSQL(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	
	script := "INSERT INTO customer(id, name, email, balance, rating, birth_date, married, created_at) VALUES('1', 'Dhandy', 'dhandy@gmail.com', 1000, 5.0, '2007-03-10', FALSE, CURRENT_TIMESTAMP)"
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
	
	script := "SELECT * FROM customer"
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