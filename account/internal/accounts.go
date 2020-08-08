package internal

import (
	"database/sql"
	"log"

	"github.com/google/uuid"
)

func CreateAccount(db *sql.DB, userID string) (string, error) {

	guid := uuid.New()
	id := guid.String()

	stmt := `insert into accounts(id, user_id)
			values($1, $2)`

	_, err := db.Exec(stmt, id, userID)
	if err != nil {
		log.Println("Error inserting user into db", err)

		return "", err
	}

	return id, nil
}
