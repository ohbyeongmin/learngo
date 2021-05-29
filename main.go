package main

import (
	"fmt"

	"github.com/ohbyeongmin/learngo/accounts"
)

func main(){
	account := accounts.NewAccount("obm")
	account.Deposit(10)
	
	fmt.Println(account)
}