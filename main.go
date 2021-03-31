package main

import "log"

func main() {
	myVar := "dog"

	switch myVar {
	case "cat":
		log.Println("myVar is set to cat")

	case "dog":
		log.Println("myVar is set to dog")

	case "fish":
		log.Println("myVar is set to fish")

	default:
		log.Println("myVar is set to something else")
	}

}
