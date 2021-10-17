package db

import (
	"database/sql"
	"fmt"
)

const (
	MAIN_DB = "postgres"
	SSL     = "disable"
)

func ConnectDB(dbName string) (*sql.DB, error) {
	connStr := fmt.Sprintf("dbname=%s sslmode=%s", dbName, SSL)
	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func CreateDB(dbName string) error {
	conn, err := ConnectDB(MAIN_DB)
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
