package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func main() {
	for {

		// Enter a password and generate a salted hash
		pwd := getPwd("password1")
		hash := hashAndSalt(pwd)

		// Enter the same password again and compare it with the
		// first password entered
		pwd2 := getPwd("password2")
		pwdMatch := comparePasswords(hash, pwd2)
		fmt.Println("Passwords Match?", pwdMatch)
	}
}
func getPwd(s string) string {
	// Prompt the user to enter a password
	fmt.Println("Enter ", s)
	// We will use this to store the users input
	var pwd string
	// Read the users input
	_, err := fmt.Scan(&pwd)
	if err != nil {
		log.Println(err)
	}
	// Return the users input as a byte slice which will save us
	// from having to do this conversion later on
	return pwd
}

func hashAndSalt(plainPwd string) string {

	// Use GenerateFromPassword to hash & salt pwd
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	hash, err := bcrypt.GenerateFromPassword([]byte(plainPwd), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
	}
	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash)
}

func comparePasswords(hashedPwd string, plainPwd string) bool {
	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, []byte(plainPwd))
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}
