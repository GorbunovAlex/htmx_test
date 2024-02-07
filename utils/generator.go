package utils

import (
	"fmt"

	"github.com/bxcodec/faker"
)

type User struct {
	FirstName string `faker:"first_name" json:"first_name" binding:"required"`
	LastName  string `faker:"last_name" json:"last_name" binding:"required"`
	Phone     string `faker:"e_164_phone_number" json:"phone" binding:"required"`
}

func GenerateRandomUsers(limit int) []User {
	users := make([]User, limit)

	for i := 0; i < limit; i++ {
		var user User
		err := faker.FakeData(&user)
		if err != nil {
			fmt.Println(err)
		}
		users[i] = user
	}

	return users
}
