package db

import (
	"database/sql"
	"fmt"
)

const (
	SSL = "disable" // For development only.
)

func Connect(driverName string, dbName string) (*sql.DB, error) {
	connStr := fmt.Sprintf("dbname=%s sslmode=%s", dbName, SSL)
	conn, err := sql.Open(driverName, connStr)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

// Create database with default database connection.
func Create(driverName string, defaultDbName string, dbName string) error {
	conn, err := Connect(driverName, defaultDbName)
	if err != nil {
		return err
	}
	conn.Exec(fmt.Sprintf("CREATE DATABASE %s", dbName))
	conn.Close()
	return nil
}

func CreateTable(conn *sql.DB, tableName string, tableSchema string) error {
	tableQuery := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (%s)", tableName, tableSchema)
	if _, err := conn.Exec(tableQuery); err != nil {
		return err
	}
	return nil
}

func Insert(conn *sql.DB, tableName string, value string) (int, error) {
	var id int
	insertQuery := fmt.Sprintf("insert into %s values (%s) RETURNING id", tableName, value)
	if err := conn.QueryRow(insertQuery).Scan(&id); err != nil {
		return -1, err
	}
	return id, nil
}
