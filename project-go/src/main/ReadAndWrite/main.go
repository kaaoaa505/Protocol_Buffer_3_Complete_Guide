package main

import (
	"io"
	"io/ioutil"
	"log"
	"strings"

	"github.com/golang/protobuf/jsonpb"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/runtime/protoiface"
	"local.dev/m/v2/build/local"
)

func main() {
	println("\n\n\n################# Read And Write - START.")

	messageToWrite := call_sample_message()
	writeToFile("message.bin", messageToWrite)

	messageToRead := &local.SampleObj{}
	data := readFromFile("message.bin", messageToRead)

	println("Data bytes array from read file is: ", data)

	println("\n---- Note: you need to use marshal to get the complete correct data")
	println("\tData from byte to string: ", string(data[:]))

	toJsonResult, toJsonError := protoToJson(messageToRead)

	if toJsonError != nil {
		println(toJsonError)
		log.Fatalln(toJsonError)
	}

	println("messageToRead protoToJson: ", toJsonResult)

	messageToProto := &local.SampleObj{}
	toProtoResult := jsonToProto(strings.NewReader(toJsonResult), messageToProto)
	println("messageToRead toProtoResult: ", toProtoResult)
	println("messageToRead messageToProto: ", messageToProto)
	println("messageToRead messageToProto.Id: ", messageToProto.Id)
	println("messageToRead messageToProto.Name: ", messageToProto.Name)
	println("messageToRead messageToProto.IsValid: ", messageToProto.IsValid)


	println("################# Read And Write - END.\n\n\n")
}

func call_sample_message() *local.SampleObj {
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

func writeToFile(fileName string, message proto.Message) error {
	println(fileName)
	output, messageError := proto.Marshal(message)

	if messageError != nil {
		println(messageError)
		log.Fatalln(messageError)
		return messageError
	}

	println(output)

	ioWriteError := ioutil.WriteFile(fileName, output, 0644)
	if ioWriteError != nil {
		println(ioWriteError)
		log.Fatalln(ioWriteError)
		return ioWriteError
	}

	println("Write to file was successful, File Name: ", fileName)

	return nil
}

func readFromFile(fileName string, message proto.Message) []byte {

	data, ioReadError := ioutil.ReadFile(fileName)

	if ioReadError != nil {
		println(ioReadError)
		log.Fatalln(ioReadError)
		return nil
	}

	println("Read from file was successful, File Name: ", fileName)

	proto.Unmarshal(data, message)

	return data
}

func protoToJson(message protoiface.MessageV1) (string, error) {
	marshaler := jsonpb.Marshaler{
		EnumsAsInts:  false,
		EmitDefaults: true,
		Indent:       "  ",
		OrigName:     true,
	}

	return marshaler.MarshalToString(message)
}

func jsonToProto(data io.Reader, message protoiface.MessageV1) error{
	unmarshalResult := jsonpb.Unmarshal(data, message)

	return unmarshalResult
}