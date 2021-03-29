package main

import "log"

type myStruct struct {
	FirstName string
}

func (m *myStruct) printFirstname() string {
	return m.FirstName
}

func main() {
	var myVar myStruct
	myVar.FirstName = "John"

	myVar2 := myStruct{
		FirstName: "Mary",
	}

	log.Println("myVar is set to", myVar.printFirstname())
	log.Println("myVar is set to", myVar2.printFirstname())
}
