package main

func main() {

	// Values of type user can be used to call methods
	// declared with both value and pointer receivers.
	bill := User{name:"Bill", email: "bill@gmail.com"}
	bill.changeEmail("b@b")
	bill.notify()

	// Pointers of type user can also be used to call methods
	// declared with both value and pointer receiver.
	joan := &User{"Joan", "joan@email.com"}
	joan.changeEmail("joan@hotmail.com")
	joan.notify()

	// Create a slice of User values with two users.
	users := []User{
		{"ed", "ed@email.com"},
		{"erick", "erick@email.com"},
	}

	// Iterate over the slice of users switching
	// semantics. Not Good!
	for _, u := range users {
		u.changeEmail("it@wontmatter.com")
	}

	// Exception example: Using pointer semantics
	// for a collection of strings.
	keys := make([]string, 10)
	for i := range keys {
		keys[i] = func() string { return "key-gen" }()
	}
}
