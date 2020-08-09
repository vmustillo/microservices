package internal

import (
	"crypto/rand"
	"database/sql"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/vmustillo/microservices/card/gen"
)

func CreateCard(db *sql.DB, id string, accountID string) (string, error) {

	cardNumExists := true
	var cardNumber string

	for cardNumExists {
		cardNumber = generateCardNumber()
		log.Println("Card Number:", cardNumber)

		var queryNum string
		checkStmt := `select card_number from cards where card_number = $1`
		err := db.QueryRow(checkStmt, cardNumber).Scan(&queryNum)
		if err != nil {
			if err != sql.ErrNoRows {
				log.Println("Error querying card number:", err)
			} else {
				cardNumExists = false
			}
		}
	}

	stmt := `insert into cards(id, card_number, account_id)
			values($1, $2, $3)`

	_, err := db.Exec(stmt, id, cardNumber, accountID)
	if err != nil {
		log.Println("Error inserting user into db", err)

		return "", err
	}

	return id, nil
}

func generateCardNumber() string {
	firstFour := "4321"

	var b strings.Builder
	b.Grow(12)

	var random *big.Int
	for i := 0; i < 12; i++ {
		random, _ = rand.Int(rand.Reader, big.NewInt(10))
		fmt.Fprintf(&b, random.String())
	}

	log.Println(b.String())

	return firstFour + b.String()
}

func GetCardFromDB(db *sql.DB, number string) (*gen.Card, error) {
	var card gen.Card

	checkStmt := `select id, card_number, account_id from cards where card_number = $1`
	err := db.QueryRow(checkStmt, number).Scan(&card.Id, &card.Number, &card.AccountId)
	if err != nil {
		log.Println("Unable to find record with card number:", number)
		return &card, err
	}

	log.Println("Card found:", card)

	return &card, nil
}
