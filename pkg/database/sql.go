package database

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type SqlClient struct {
	db *sql.DB
}

func NewSqlClient(ctx context.Context, dbType string, host string, port string, username string, password string, dbName string) (*SqlClient, error) {

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", username, password, host, port, dbName)
	if dbType == MSSQL {
		connectionString = fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s;", host, username, password, port, dbName)
	}

	db, err := sql.Open(dbType, connectionString)
	if err != nil {
		return nil, fmt.Errorf("failed to create sql connection: %w", err)
	}

	err = db.PingContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to sql: %w", err)
	}

	return &SqlClient{db: db}, nil
}

func (sql *SqlClient) GetDatabaseDetails() string {

	var result string

	// Run query and scan for result
	err := sql.db.QueryRowContext(context.Background(), "SELECT @@version").Scan(&result)

	if err != nil {
		return fmt.Sprintf("Database scan failed: %v", err.Error())
	}

	returnStr := strings.Split(strings.ReplaceAll(result, "\t", ""), "\n")

	driverType, _, _ := strings.Cut(reflect.TypeOf(sql.db.Driver()).String(), ".")

	if strings.Contains(driverType, MSSQL) {
		return returnStr[:len(returnStr)-1][0]
	}

	if strings.Contains(driverType, MYSQL) {
		return returnStr[0]
	}

	return "No return definition this sql driver. Please contact an administrator."

}

func (sql *SqlClient) GetDB() *sql.DB {
	return sql.db
}
