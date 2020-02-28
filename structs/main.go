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

type bill struct {
	flag 	bool
	counter int16
	pi 		float32
}

type alice struct {
	flag 	bool
	counter int16
	pi 		float32
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

	//Declare avariable of an anonymous type set its zero value.
	var e1 struct{
		flag 	bool
		counter int16
		pi		float32
	}

	e2 := struct {
		flag 	bool
		counter int16
		pi 		float32
	}{
		flag:    false,
		counter: 1,
		pi:      1,
	}

	fmt.Printf("%+v\n", e1)
	fmt.Printf("%+v\n", e2)

	var a alice
	var b bill

	//implicit conversion (not alove)
	// a = b

	//exlicitly convert
	b = bill(a)

	b = e2

	fmt.Printf("%+v\n", a)
	fmt.Printf("%+v\n", b)


}
