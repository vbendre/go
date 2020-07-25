package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	contact   ContactInfo
}

type ContactInfo struct {
	email   string
	zipcode int
}

func (p Person) print() {
	fmt.Printf("%+v", p)
}

func (p Person) passByValueUpdateName(fName string, lName string) {
	p.firstName = fName
	p.lastName = lName
}

func (pointerPerson *Person) passByReferenceUpdateName(fName string, lName string) {

	(*pointerPerson).firstName = fName // derefering the pointer
	(*pointerPerson).lastName = lName

}

func main() {

	johnContact := ContactInfo{email: "john.doe@gmail.com", zipcode: 11111}
	john := Person{firstName: "john", lastName: "doe", contact: johnContact}

	john.firstName = "john walker"

	// var test Person
	// test.print()

	john.passByValueUpdateName("john1", "walker1")
	john.print()

	//explicit pointer conversion
	johnPointer := &john
	johnPointer.passByReferenceUpdateName("john1", "walker1")
	john.print()

	//implicit pointer conversion
	john.passByReferenceUpdateName("john2", "walker 2")
	john.print()

	mySlice := []string{"hi", "there"}
	//slices are passed by reference by default unlike struct
	updateSlice(mySlice)

	fmt.Println(mySlice)

}

func updateSlice(argSlice []string) {
	argSlice[0] = "bye"
}
