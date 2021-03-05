package main

import (
	"github.com/brianvoe/gofakeit/v6"
)

type user struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Address    string `json:"address"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
	PostalCode string `json:"postalCode"`
	City       string `json:"city"`
	Number     string `json:"number"`
	ID         string `json:"id"`
}

func generateUsers(n int) []user {
	users := make([]user, 0, n)

	for i := 0; i < n; i++ {

		newUser := user{
			Name:       gofakeit.FirstName(),
			Surname:    gofakeit.LastName(),
			Address:    gofakeit.Street(),
			Phone:      gofakeit.Phone(),
			Email:      gofakeit.Email(),
			PostalCode: gofakeit.Zip(),
			City:       gofakeit.City(),
			Number:     gofakeit.StreetNumber(),
			ID:         gofakeit.UUID(),
		}

		users = append(users, newUser)

	}

	return users
}
