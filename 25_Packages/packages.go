package main

import (
	"fmt"
	"module_example/25_Packages/auth"
	"module_example/25_Packages/user"
)

func main() {
	auth.LoginWithCredentials("Nazeer", "secret123")
	session := auth.GetSession()
	fmt.Println("session", session)

	user:= user.User{
		Email: "user@email.com",
		Name:  "Nazeer",
	}
	fmt.Println(user.Email, user.Name)
}
