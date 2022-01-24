package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func main() {
	// password := "Sumathy"
	// hash, _ := HashPassword(password) // ignore error for the sake of simplicity

	var password string
	fmt.Println("Enter Password to encrypt: ")
	fmt.Scanf("%s", &password)

	hash, _ := HashPassword(password)

	fmt.Println("Password Entered:", password)
	fmt.Println("Hash Value:    ", hash)

	match := CheckPasswordHash(password, hash)
	fmt.Println("Match Found:   ", match)
}
