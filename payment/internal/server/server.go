package server

import (
	"context"
	"log"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/google/uuid"
	account "github.com/vmustillo/microservices/account/gen"
	card "github.com/vmustillo/microservices/card/gen"
	"github.com/vmustillo/microservices/payment/gen"
)

func (s *PaymentServer) MakePayment(ctx context.Context, req *gen.MakePaymentRequest) (*gen.Payment, error) {
	log.Println("Request", req)
	var wg sync.WaitGroup

	cardChan := make(chan bool)
	wg.Add(1)
	go cardNumberExists(req.GetCardNumber(), cardChan, &wg)

	cardExists := <-cardChan
	log.Println("CardExists:", cardExists)

	accountChan := make(chan bool)
	wg.Add(1)
	go isBelowBalance(req.GetPurchaserAccountId(), req.GetAmount(), accountChan, &wg)

	belowBal := <-accountChan
	wg.Wait()
	log.Println("Below Bal:", belowBal)

	if !cardExists || !belowBal {
		log.Println("Either Card Does not exist or balance is too low. Cancelling transaction")
		return &gen.Payment{}, nil
	}

	guid := uuid.New()
	id := guid.String()

	return &gen.Payment{Id: id, PurchaserAccountId: req.GetPurchaserAccountId(), PurchasedFromAccountId: req.GetPurchasedFromAccountId(), Amount: req.GetAmount()}, nil
}

func NewPaymentServer(paySrv *PaymentServer) *grpc.Server {

	srv := grpc.NewServer()
	gen.RegisterPaymentServiceServer(srv, paySrv)
	reflection.Register(srv)

	return srv
}

func cardNumberExists(cardNum string, cardChan chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	conn, err := grpc.Dial("localhost:4042", grpc.WithInsecure())
	if err != nil {
		log.Println("Card Request Done: False")
		cardChan <- false
		return
	}

	cardConn := card.NewCardServiceClient(conn)
	log.Println("Payment card number:", cardNum)
	res, err := cardConn.GetCard(context.Background(), &card.GetCardRequest{Number: cardNum})
	if err != nil {
		log.Println("Card Request Done: False")
		log.Println("Error requesting card service", err)
		cardChan <- false
		return
	}
	if res == nil {
		log.Println("Card Request Done: False")
		cardChan <- false
		return
	}

	log.Println("Card Request Done: True")
	cardChan <- true
}

func isBelowBalance(accountID string, amount float64, accountChan chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err != nil {
		log.Println("Card Request Done: False")
		accountChan <- false
		return
	}

	accountConn := account.NewAccountServiceClient(conn)
	log.Println("Account ID:", accountID)
	res, err := accountConn.GetAccount(context.Background(), &account.GetAccountRequest{Id: accountID})
	if err != nil {
		log.Println("Account Request Done: False")
		log.Println("Error requesting account service", err)
		accountChan <- false
		return
	}
	if res == nil {
		log.Println("Account Request Done: False")
		accountChan <- false
		return
	}

	bal := res.Balance
	limit := res.CreditLimit
	log.Println("Balance:", bal, "Ammount:", amount)
	log.Println("Limit:", limit)
	if bal+amount > limit {
		log.Println("Account Request Done: False")
		log.Println("Over credit limit: ")
		accountChan <- false
		return
	}

	log.Println("Account Request Done: True")
	accountChan <- true
}
