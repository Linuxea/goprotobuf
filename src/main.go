package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"os"
	"test"
)

func write() {
	p1 := &test.Person{
		Id:   1,
		Name: "little zhang",
		Phones: []*test.Phone{
			{Type: test.PhoneType_HOME, Number: "1111111111"},
			{Type: test.PhoneType_WORK, Number: "22222222222"},
		},
	}

	p2 := &test.Person{
		Id:   2,
		Name: "little wang",
		Phones: []*test.Phone{
			{Type: test.PhoneType_HOME, Number: "333333333"},
			{Type: test.PhoneType_WORK, Number: "4444444444"},
		},
	}

	// create address book
	book := &test.ContactBook{}
	book.Persons = append(book.Persons, p1)
	book.Persons = append(book.Persons, p2)

	// encode
	data, _ := proto.Marshal(book)
	// write data to file
	ioutil.WriteFile("./test", data, os.ModePerm)
}

func read() {

	data, _ := ioutil.ReadFile("./test")
	book := &test.ContactBook{}
	// decode data
	proto.Unmarshal(data, book)
	for _, v := range book.Persons {
		fmt.Println(v.Id, v.Name)
		for _, vv := range v.Phones {
			fmt.Println(vv.Type, vv.Number)
		}
	}

}

func main() {
	write()
	read()

}
