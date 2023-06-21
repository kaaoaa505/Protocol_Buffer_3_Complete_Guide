package main

import (
	"io/ioutil"
	"log"
	"strings"

	"google.golang.org/protobuf/proto"
	"local.dev/m/v2/build/local"
)

func main() {
	println("Hello World!.")

	message := call_sample_message()

	println(message)

	writeToFile("message.bin", message);
}

func call_sample_message() *local.SampleObj{
	sm := local.SampleObj{
		Id:      0,
		Name:    "",
		IsValid: false,
		StrList: []string{"str1", "str2", "str3"},
	}

	sm.IsValid = true
	sm.Id = 123
	sm.Name = "Khaled Allam"

	println(sm.IsValid)
	println(sm.Id)
	println(sm.Name)
	println(strings.Join(sm.StrList, ","))

	return &sm
}

func writeToFile(fileName string, message proto.Message) error{
	println(fileName)
	output, messageError := proto.Marshal(message);

	if(messageError != nil){
		println(messageError)
		log.Fatalln(messageError);
		return messageError;
	}

	println(output)

	ioWriteError := ioutil.WriteFile(fileName, output, 0644);
	if(ioWriteError != nil){
		println(ioWriteError)
		log.Fatalln(ioWriteError);
		return ioWriteError;
	}

	println("Write to file was successful, File Name: ", fileName)

	return nil;
}