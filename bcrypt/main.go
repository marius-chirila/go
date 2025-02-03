package main

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	password := `123456`
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(hashedPassword)
	loginPwd := `1233456`
	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(loginPwd))
	if err == nil {
		fmt.Println("Password successful")
	} else {
		log.Fatalf("Provided pwd not the same.%v", err)
	}
}
