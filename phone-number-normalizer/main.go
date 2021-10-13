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
	if len(os.Args) <= 1 {
		log.Println("Please supply a single word.")
		return
	}
	word := os.Args[1]

	connStr := "dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
		return
	}

	_, err = db.Exec(fmt.Sprintf("insert into users values ('%s')", word))
	if err != nil {
		log.Fatal(err)
		return
	}

	if rows, err := db.Query("SELECT * FROM users"); err == nil {
		for rows.Next() {
			var (
				name string
			)
			if err := rows.Scan(&name); err != nil {
				log.Fatal(err)
			}
			println(name)
		}
	} else {
		log.Fatal("Query failed, reason: ", err)
	}
}
