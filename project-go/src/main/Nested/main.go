package main

import (
	"fmt"

	"local.dev/m/v2/build/local"
)

func main() {
	println("\n\n\n################# Nested Main - START.")

	message := local.NestedObj{}

	message.BirthDate = &local.NestedObj_Date{
		Year: 1988,
	}

	message.EyeColor = local.NestedObj_BROWN

	firstAddress := local.NestedObj_Address{
		AddressLine_1: "123 Fake st.",
		AddressLine_2: "Fake City",
		City: "Any City",
		Country: "Any Country",
	}

	message.Addresses = append(message.Addresses, &firstAddress)


	// println(message.Addresses)
	fmt.Println(message.Addresses)

	println(message.Addresses[len(message.Addresses)-1].AddressLine_1)
	println(message.Addresses[len(message.Addresses)-1].AddressLine_2)
	println(message.Addresses[len(message.Addresses)-1].City)
	println(message.Addresses[len(message.Addresses)-1].Country)
	println(message.EyeColor.String())

	// println(message)
	fmt.Println(message.String())

	println("################# Nested Main - END.\n\n\n")
}

