package main

import "fmt"

type contact struct {
	email   string
	zipCode int
}
type person struct {
	firstName   string
	lastName    string
	contactInfo contact
}

func (p person) print() {
	fmt.Println(p)
}

func (p *person) updateName(name string) {
	(*p).firstName = name
}

func main() {

	p := person{
		firstName: "Stefan",
		lastName:  "Halapa",
		contactInfo: contact{
			email:   "halapa",
			zipCode: 123}}

	var stefan person

	stefan.lastName = "stefan"
	stefan.firstName = "halapa"

	fmt.Println(stefan)
	fmt.Println(p)
	fmt.Println(p.firstName)
	fmt.Printf("%+v", p)

	stefan.firstName = "Stt"

	fmt.Println("-----------------------")
	stefan.print()

	stefanPointer := &stefan

	stefanPointer.updateName("Pointer")
	stefanPointer.print()
	stefan.updateName("booooo")
	stefan.print()

	fmt.Println("--------------------------------------")

}
