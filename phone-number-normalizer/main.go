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
	MAIN_DB    = "postgres"
	DB_NAME    = "phone_directory"
	SSL        = "disable"
	MAIN_TABLE = "phone_numbers"
)

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

func scream(err error) {
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func main() {
	scream(createDB(DB_NAME))
	conn, err := connectDB(DB_NAME) // Use the newly created database
	scream(err)
	scream(createTable(conn, "phone_numbers", `
		id SERIAL PRIMARY KEY,
		number VARCHAR(255)
	`))
}
