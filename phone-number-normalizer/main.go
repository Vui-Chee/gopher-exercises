package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"

	_ "github.com/lib/pq"
)

const (
	MAIN_DB     = "postgres"
	PHONE_TABLE = "phone_numbers"
	DB_NAME     = "phone_directory"
	SSL         = "disable"
)

func scream(err error) {
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func normalize(phone string) string {
	var builder strings.Builder
	for _, ch := range phone {
		if unicode.IsDigit(ch) {
			builder.WriteRune(ch)
		}
	}
	return builder.String()
}

func connectDB(dbName string) (*sql.DB, error) {
	connStr := fmt.Sprintf("dbname=%s sslmode=%s", dbName, SSL)
	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func createDB(dbName string) error {
	conn, err := connectDB(MAIN_DB)
	if err != nil {
		return err
	}
	conn.Exec(fmt.Sprintf("CREATE DATABASE %s", dbName))
	conn.Close()
	return nil
}

func createTable(conn *sql.DB, tableName string, tableSchema string) error {
	tableQuery := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (%s)", tableName, tableSchema)
	if _, err := conn.Exec(tableQuery); err != nil {
		return err
	}
	return nil
}

func insertSingleEntry(conn *sql.DB, tableName string, value string) (int, error) {
	var id int
	insertQuery := fmt.Sprintf("insert into %s values (%s) RETURNING id", tableName, value)
	if err := conn.QueryRow(insertQuery).Scan(&id); err != nil {
		return -1, err
	}
	return id, nil
}

func main() {
	phone_numbers := []string{
		"1234567890",
		"123 456 7891",
		"(123) 456 7892",
		"(123) 456-7893",
		"123-456-7894",
		"(123)456-7892",
	}

	scream(createDB(DB_NAME))
	conn, err := connectDB(DB_NAME) // Use the newly created database
	scream(err)
	scream(createTable(conn, PHONE_TABLE, `
		id SERIAL PRIMARY KEY,
		number VARCHAR(255)
	`))

	for _, phone_number := range phone_numbers {
		rowEntry := fmt.Sprintf("default, '%s'", normalize(phone_number))
		id, err := insertSingleEntry(conn, PHONE_TABLE, rowEntry)
		scream(err)
		fmt.Printf("Inserted id %d\n", id)
	}
}
