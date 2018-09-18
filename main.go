package main

import (
	"fmt"
)

var wallets []wallet

func main() {
	initWallets()
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
			balance: 0,
		},
	}
}

func showWallets() {
	for _, item := range wallets {
		fmt.Printf("Wallet %[1]d has balance %s\n", item.id, item.getBalanceString())
	}
}
