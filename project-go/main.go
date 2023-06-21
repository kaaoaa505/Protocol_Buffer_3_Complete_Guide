package main

import (
	"fmt"
	"local.dev/m/v2/build/local"
)

func main() {
	fmt.Println("Hello World!.")

	call_sample_message()
}

func call_sample_message() {
	sm := local.SampleObj{
		Id:      0,
		Name:    "",
		IsValid: false,
		StrList: []string{"str1", "str2", "str3"},
	}

	sm.IsValid = true
	sm.Id = 123
	sm.Name = "Khaled Allam"

	fmt.Println(sm.IsValid)
	fmt.Println(sm.Id)
	fmt.Println(sm.Name)
	fmt.Println(sm.StrList)
}
