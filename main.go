package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var wallets []wallet

type addRequestDto struct {
	Balance int
}

type getBalanceRequestDto struct {
	ID int
}

type makePaymentRequestDto struct {
	SenderID    int
	RecipientID int
	Sum         uint
}

func main() {
	initWallets()

	http.HandleFunc("/api/add", add)
	http.HandleFunc("/api/getBalance", getBalance)
	http.HandleFunc("/api/makePayment", makePayment)
	http.HandleFunc("/api/showAll", showAll)

	http.ListenAndServe(":8181", nil)
}

func add(rw http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		return
	}

	decoder := json.NewDecoder(req.Body)

	var reqDto addRequestDto
	errParse := decoder.Decode(&reqDto)

	if errParse != nil {
		fmt.Println(errParse)
		return
	}

	id := addWallet(reqDto.Balance)

	fmt.Println(id)
}

func getBalance(rw http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		return
	}

	decoder := json.NewDecoder(req.Body)

	var reqDto getBalanceRequestDto
	errParse := decoder.Decode(&reqDto)

	if errParse != nil {
		fmt.Println(errParse)
	}

	wal, errGet := getWallet(reqDto.ID)
	if errGet != nil {
		fmt.Println(errGet)
		return
	}

	fmt.Println(wal.getBalanceString())
}

func makePayment(rw http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		return
	}

	decoder := json.NewDecoder(req.Body)

	var reqDto makePaymentRequestDto
	errParse := decoder.Decode(&reqDto)

	if errParse != nil {
		fmt.Println(errParse)
		return
	}

	senderWallet, errGetSender := getWallet(reqDto.SenderID)
	if errGetSender != nil {
		fmt.Println(errGetSender)
		return
	}

	recipientWallet, errGetRecipient := getWallet(reqDto.RecipientID)
	if errGetRecipient != nil {
		fmt.Println(errGetRecipient)
		return
	}

	senderWallet.makePayment(recipientWallet, reqDto.Sum)
}

func showAll(rw http.ResponseWriter, req *http.Request) {
	showWallets()
}

func initWallets() {
	wallets = []wallet{
		wallet{
			id:      1,
			balance: 0,
		},
		wallet{
			id:      2,
			balance: 2000,
		},
	}
}

func getWallet(id int) (*wallet, error) {
	for index, item := range wallets {
		if item.id == id {
			return &wallets[index], nil
		}
	}

	var result wallet
	return &result, fmt.Errorf("Wallet with id %d doesn't exist", id)
}

func addWallet(initBalance int) int {
	newWallet := wallet{
		balance: initBalance,
	}
	if len(wallets) <= 0 {
		newWallet.id = 1
	} else {
		newWallet.id = wallets[len(wallets)-1].id + 1
	}
	wallets = append(wallets, newWallet)
	return newWallet.id
}

func showWallets() {
	for _, item := range wallets {
		fmt.Printf("Wallet %[1]d has balance %s\n", item.id, item.getBalanceString())
	}
}
