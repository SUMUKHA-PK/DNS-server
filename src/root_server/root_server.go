package root_server

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

var receive string

func Root_server() {

	// Starting the server
	fmt.Printf("I am the root server!\n")
	link, err := net.Listen("tcp", ":12345")
	if err != nil {
		fmt.Print(err)
		fmt.Printf("\n2\n")
	}

	//Continous server listening
	for {

		fmt.Printf("\nServer listening for incoming connections on port 12345\n\n")
		conn, err := link.Accept()
		if err != nil {
			fmt.Print(err)
			fmt.Printf("\n1\n")
		}

		//Create a new scanner and get the data from the client
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			receive = scanner.Text()
			fmt.Printf("IP received to map from client: " + receive + "\n")

			break
		}
		if errReadConn := scanner.Err(); errReadConn != nil {
			fmt.Print(errReadConn)
			return
		}

		//Get the result of the mapping from the servers
		result := start_servers(receive)

		//Communicate back the result to the client on the same connection
		scanner = bufio.NewScanner(strings.NewReader(result))

		for scanner.Scan() {
			text := scanner.Text()
			_, err := fmt.Fprintf(conn, text+"\n")
			if err != nil {
				//Error exists due to sending in same connection, figure it out
			}
			log.Print("Query mapping sent: " + text)
			break
		}
	}
}

func start_servers(IP string) string {
	log.Printf("Starting servers, sending data: " + IP + "\n")

	split := strings.Split(IP, ".")
	fmt.Printf("Length of split: %d\n", len(split))

	return "bitch"
}
