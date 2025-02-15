package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type User struct {
	FirstName string `validate:"required"`
	LastName  string `validate:"required"`
	Email     string `validate:"required,email"`
	Age       int    `validate:"gte=0,lte=130"`
}

func validateUser(user User) error {
	validate := validator.New()
	return validate.Struct(user)
}

func test() {
	user := User{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@example.com",
		Age:       30,
	}

	err := validateUser(user)
	if err != nil {
		fmt.Println("Validation failed:", err)
	} else {
		fmt.Println("Validation succeeded")
	}
}

func multiply(a int, b int) int {
	println("Multiplication of 10 and 20 is: ", a*b)
	return a * b
}
