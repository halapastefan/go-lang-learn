package main

import "fmt"

type User struct {
	name 	string
	email 	string
}

func (u User) notify() {
	fmt.Printf("Sending user email to %s<<%s>>\n", u.name, u.email)
}

func (u *User) changeEmail(email string) {
	u.email = email
}