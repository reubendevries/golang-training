package main

import (
	"log"
	"time"
)

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

func main() {
	user := User{
		FirstName:   "Reuben",
		LastName:    "deVries",
		PhoneNumber: "1 555 555-1212",
	}

	log.Println(user.FirstName, user.LastName, "Birthdate:", user.BirthDate)
}
