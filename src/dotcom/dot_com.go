package dotcom

import (
	"fmt"
	"net"
)

func DotCom() {
	fmt.Printf("I am the .com server!\n")

	// Starting the server
	link, err := net.Listen("tcp", "127.0.1.1:12345")
	if err != nil {
		fmt.Print(err)
		fmt.Printf("\n2\n")
	}

}
