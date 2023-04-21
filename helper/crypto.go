package helper

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

const salt = 5

func HashPassword(pw string) string {
	hashed, err := bcrypt.GenerateFromPassword([]byte(pw), salt)
	if err != nil {
		log.Fatal(err)
	}
	return string(hashed)
}

func ComparePassword(hashed, pw string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(pw))
	if err != nil {
		return false
	}
	return true
}
