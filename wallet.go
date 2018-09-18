package main

import (
	"fmt"
	"strconv"
)

type iWallet interface {
	addBalance(addedSum int)
	makePayment(recipient *wallet, paymentSum uint)
}

type wallet struct {
	id      int
	balance int
}

func (currentWallet *wallet) addBalance(addedSum int) wallet {
	(*currentWallet).balance += addedSum
	return *currentWallet
}

func (currentWallet *wallet) makePayment(recipientWallet *wallet, paymentSum uint) bool {
	if currentWallet.balance <= int(paymentSum) {
		fmt.Println(fmt.Sprintf("Нужно больше золота! Доступно %[1]s, запрошено %[2]s", currentWallet.getBalanceString(), getStringFromSum(int(paymentSum))))
		return false
	}

	(*currentWallet).balance -= int(paymentSum)
	(*recipientWallet).balance += int(paymentSum)
	fmt.Println(fmt.Sprintf("Перевод с кошелька %[1]d на кошелек %[2]d суммы %[3]s прошел успешно", currentWallet.id, recipientWallet.id, getStringFromSum(int(paymentSum))))

	return true
}

func (currentWallet wallet) getBalanceString() (result string) {
	return getStringFromSum(currentWallet.balance)
}

func getStringFromSum(sum int) (result string) {
	result = strconv.Itoa(sum/100) + ","
	remainder := sum % 100
	result += strconv.Itoa(remainder)
	if remainder == 0 {
		result += "0"
	}
	return
}
