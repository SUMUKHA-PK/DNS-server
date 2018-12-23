package root_server

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func Root_server(IP []string) {
	fmt.Printf("I am the root server!\n")
	fmt.Print(IP)
	link, err := net.Listen("tcp", ":12345")
	if err != nil {
		fmt.Print(err)
	}
	for {
		fmt.Printf("\nServer listening for incoming connections on port 12345\n")
		conn, err := link.Accept()
		if err != nil {
			fmt.Print(err)
		}
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			fmt.Printf("Client sends: " + scanner.Text())
			break
		}
		if errReadConn := scanner.Err(); errReadConn != nil {
			log.Printf("Read error: %T %+v", errReadConn, errReadConn)
			return
		}
	}
}
