package main

import (
	"fmt"
	"log"

	"github.com/golang/protobuf/proto"
)

func main() {

	elp := &Person{
		Name: "Avinash Maurya",
		Age:  24,
	}

	data, err := proto.Marshal(elp)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	// printing out our raw protobuf object
	fmt.Println(data)

	newElp := &Person{}
	err = proto.Unmarshal(data, newElp)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}
	fmt.Println(newElp.GetName())
	fmt.Println(newElp.GetAge())

}
