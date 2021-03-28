package main

import "fmt"

func main() {
	// Define map
	//emails := make(map[string]string)

	// Assign keyvalue pairs
	// emails["Bob"] = "bob@gmail.com"
	// emails["Sharon"] = "sharon@gmail.com"
	// emails["danthe"] = "danthe@gmail.com"

	//Declare map and add keyvalue pairs
	emails := map[string]string{"Bob": "bob@gmail.com", "danthe": "danthe@gmail.com"}

	// Can add like this anyway
	emails["Sharon"] = "sharon@gmail.com"

	fmt.Println(emails)
	fmt.Println(emails["Bob"])
	fmt.Println(len(emails))

	// Delete from map
	delete(emails, "Bob")
	fmt.Println(emails)

}
