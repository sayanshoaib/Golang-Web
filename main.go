package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	password := "password"
	pass, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		panic(err)
	}
	fmt.Println("Main Password", password)
	fmt.Println("Hashed Password", string(pass))

	compare := bcrypt.CompareHashAndPassword([]byte(pass), []byte(password))
	if compare == nil {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}

	pass1 := "123"

	compare2 := bcrypt.CompareHashAndPassword([]byte(pass), []byte(pass1))
	if compare2 == nil {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
}
