package main

import (
	"fmt"
	"strings"
)

func main() {

	siteName := "SuperCharacters"
	const conferenceCharacters int = 20
	var remainingCharacters uint = 20
	buyers := []string{}

	fmt.Println("Welcome to", siteName)
	fmt.Printf("We have total of %v character and %v are still available.\n", conferenceCharacters, remainingCharacters)

	for {
		var userName string
		var userLastname string
		var userEmail string
		var userAge int
		var userCharacter uint

		fmt.Println("Enter ur name:")
		fmt.Scan(&userName)

		fmt.Println("Enter ur Surname:")
		fmt.Scan(&userLastname)

		fmt.Println("Enter ur email:")
		fmt.Scan(&userEmail)

		fmt.Println("Enter ur Age:")
		fmt.Scan(&userAge)

		fmt.Println("Enter amount of character you want:")
		fmt.Scan(&userCharacter)

		remainingCharacters = remainingCharacters - userCharacter

		fmt.Printf("Thanks u %v %v for buy %v characters. You will receive a confirmation email at %v \n", userName, userLastname, userCharacter, userEmail)
		fmt.Printf("Now we have %v characters remaining for %v \n", remainingCharacters, siteName)

		buyers = append(buyers, userName)

		firstNames := []string{}
		// _ or blank identifier
		for _, buyer := range buyers {
			// Fields separates the elements with blank space
			var names = strings.Fields(buyer)
			firstNames = append(firstNames, names[0])

		}
		fmt.Printf("The first names of buyers are: %v \n", firstNames)
	}

}
