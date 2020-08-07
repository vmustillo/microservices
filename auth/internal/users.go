package internal

import (
	"database/sql"
	"log"

	"github.com/google/uuid"
	"golang.org/x/crypto/argon2"
)

func CreateUser(db *sql.DB, firstName string, lastName string, username string, password string) (string, error) {
	salt, err := GenerateSaltBytes()
	if err != nil {
		log.Println("Error generating salt bytes", err)

		return "", err
	}

	hashedPwd := argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)

	guid := uuid.New()
	id := guid.String()

	stmt := `insert into users(id, first_name, last_name, full_name, username, salt, hashed_password)
			values($1, $2, $3, $4, $5, $6, $7)`

	_, err = db.Exec(stmt, id, firstName, lastName, firstName+" "+lastName, username, salt, hashedPwd)

	if err != nil {
		log.Println("Error inserting user into db", err)

		return "", err
	}

	return id, nil
}
