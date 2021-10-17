package main

import (
	"fmt"

	_ "github.com/lib/pq"
	"github.com/vui-chee/gopher-exercises/phone-number-normalizer/db"
)

const (
	DEFAULT_DB  = "postgres"
	DRIVER      = "postgres"
	PHONE_TABLE = "phone_numbers"
	DB_NAME     = "phone_directory"
)

func main() {
	phone_numbers := []string{
		"1234567890",
		"123 456 7891",
		"(123) 456 7892",
		"(123) 456-7893",
		"123-456-7894",
		"(123)456-7892",
	}

	scream(db.Create(DRIVER, DEFAULT_DB, DB_NAME))
	conn, err := db.Connect(DRIVER, DB_NAME) // Use the newly created database
	scream(err)
	scream(db.CreateTable(conn, PHONE_TABLE, `
		id SERIAL PRIMARY KEY,
		number VARCHAR(255)
	`))

	for _, phone_number := range phone_numbers {
		rowEntry := fmt.Sprintf("default, '%s'", normalize(phone_number))
		id, err := db.Insert(conn, PHONE_TABLE, rowEntry)
		scream(err)
		fmt.Printf("Inserted id %d\n", id)
	}
}
