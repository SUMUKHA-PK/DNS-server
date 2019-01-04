package google

import (
	"fmt"
	"net"
)

func Google() {
	fmt.Printf("I am the google authoritative server!\n")

	// Starting the server
	link, err := net.Listen("tcp", "127.0.2.1:12345")
	if err != nil {
		fmt.Print(err)
		fmt.Printf("\n2\n")
	}
}
