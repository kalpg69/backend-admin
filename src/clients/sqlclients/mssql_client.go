package sqlclients

import (
	"context"
	"database/sql"
	"log"
	"net/url"

	_ "github.com/denisenkom/go-mssqldb"
)

func NewMSSQLClient(ctx context.Context, userName string, password string, serverName string, databaseName string, timeout string) (*sql.DB, error) {
	query := url.Values{}
	query.Add("database", databaseName)
	query.Add("connection+timeout", timeout)

	conn := &url.URL{
		Scheme:   "sqlserver",
		User:     url.UserPassword(userName, password),
		Path:     serverName,
		RawQuery: query.Encode(),
	}

	db, err := sql.Open("sqlserver", conn.String())
	if err != nil {
		log.Fatalf("Error: An error occurred while openning database connection : %v", err)
		return nil, err
	}

	if err := db.PingContext(ctx); err != nil {
		log.Fatalf("Error: An error occurred while connecting to database : %v", err)
		return nil, err
	}
	log.Println("Database connected!")
	return db, nil
}
