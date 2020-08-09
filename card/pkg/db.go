package pkg

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func GetConnection(dbType string, user string, endpoint string, port int, name string) (*sql.DB, error) {

	connStr := fmt.Sprintf("%s://%s@%s:%d/%s?sslmode=disable", dbType, user, endpoint, port, name)
	log.Println("Connection String: ", connStr)

	db, err := sql.Open(dbType, connStr)
	if err != nil {
		log.Println("Error opening connection to db: ", err)

		return nil, err
	}

	return db, nil
}
