package internal

import (
	"bytes"
	"crypto/rand"
	"database/sql"
	"errors"
	"log"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/argon2"
)

func GenerateSaltBytes() ([]byte, error) {
	length := 32
	saltBytes := make([]byte, length)
	_, err := rand.Read(saltBytes)
	if err != nil {
		return []byte{}, err
	}

	return saltBytes, nil
}

func Login(db *sql.DB, username string, password string) (string, error) {

	stmt := `select salt, hashed_password from users where username = $1`
	var salt []byte
	var hashedPwd []byte

	rows, err := db.Query(stmt, username)
	if err != nil {
		log.Printf("Error querying db from statement: %s\n Error: %s\n", stmt, err)

		return "", err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&salt, &hashedPwd)
		if err != nil {
			log.Printf("Error querying db from statement: %s\n Error: %s\n", stmt, err)

			return "", err
		}
	}

	hashed := argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)

	if bytes.Compare(hashed, hashedPwd) == 0 {
		log.Println("Passwords Match")
	} else {
		log.Println("Passwords do not match")

		return "", errors.New("Passwords do not match")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"group": "user",
	})

	tokenString, err := token.SignedString([]byte("SECRET_STRING"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
