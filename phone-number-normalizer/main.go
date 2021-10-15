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

var pool *sql.DB

func normalize(phone string) string {
	var builder strings.Builder

	for _, ch := range phone {
		if unicode.IsDigit(ch) {
			builder.WriteRune(ch)
		}
	}

	return builder.String()
}

func main() {
	// Get arg
	if len(os.Args) <= 1 {
		log.Println("Please supply a single word.")
		return
	}
	word := os.Args[1]

	// Try connect
	connStr := "dbname=postgres sslmode=disable"
	pool, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer pool.Close()

	// Insert db
	_, err = pool.Exec(fmt.Sprintf("insert into users values ('%s')", word))
	if err != nil {
		log.Fatal(err)
		return
	}

	// Query db
	rows, err := pool.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal("Query failed, reason: ", err)
		return
	}

	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			log.Fatal(err)
		}
		println(name)
	}
}
