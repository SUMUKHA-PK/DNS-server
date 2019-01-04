package dotedu

import (
	"fmt"
	"net"
)

func DotEdu() {
	fmt.Printf("I am the .edu server!\n")

	// Starting the server
	link, err := net.Listen("tcp", "127.0.1.2:12345")
	if err != nil {
		fmt.Print(err)
		fmt.Printf("\n2\n")
	}
}
