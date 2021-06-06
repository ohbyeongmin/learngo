package main

import (
	"fmt"

	"github.com/ohbyeongmin/learngo/mydict"
)

func main(){
	// account := accounts.NewAccount("obm")
	// account.Deposit(10)
	
	// fmt.Println(account)

	// dictionary := mydict.Dictionary{"first": "First word"}
	// definition, err := dictionary.Search("second")
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(definition)
	// }

	// dictionary := mydict.Dictionary{}
	// err := dictionary.Add("hello", "Greeting")

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// definition, err2 := dictionary.Search("hello")
	// if err2 != nil {
	// 	fmt.Println(err2)
	// }
	// fmt.Println(definition)

	// dictionary := mydict.Dictionary{}
	// baseWord := "hello"
	// dictionary.Add(baseWord, "First")
	// err := dictionary.Update(baseWord, "Second")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// word, _ := dictionary.Search(baseWord);
	// fmt.Println(word)

	dictionary := mydict.Dictionary{}
	baseWord := "hello"
	dictionary.Add(baseWord, "First")
	dictionary.Delete(baseWord)
	word, err := dictionary.Search(baseWord)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(word)
}