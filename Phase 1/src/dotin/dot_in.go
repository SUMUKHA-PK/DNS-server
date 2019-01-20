package dotin

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"

	"../helper"
)

var receive string

func DotIn(IP_List_Name []string, IP_List_Addr []string) {

	//Starting the server
	link, err := net.Listen("tcp", "127.0.1.0:12345")
	if err != nil {
		fmt.Print("Dotin server: ")
		fmt.Print(err)
	}

	//Continous server listening
	for {

		fmt.Printf("\nDotIn server listening for incoming connections on port 12345\n\n")
		conn, err := link.Accept()
		if err != nil {
			fmt.Print("Dotin server: ")
			fmt.Print(err)
		}

		//Create a new scanner and get the data from the client
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			receive = scanner.Text()
			fmt.Print("Dotin server: ")
			fmt.Printf("IP received to map from client: %s\n", receive)
			break
		}
		if errReadConn := scanner.Err(); errReadConn != nil {
			fmt.Print(errReadConn)
			return
		}

		//Get the IP from the below servers
		split := strings.Split(receive, ".")
		server_name := split[len(split)-1]
		result := helper.ContactHelper(receive, IP_List_Name, IP_List_Addr, server_name)

		//Communicate back the result to the client on the same connection
		scanner = bufio.NewScanner(strings.NewReader(result))

		for scanner.Scan() {
			text := scanner.Text()
			_, err := fmt.Fprintf(conn, text+"\n")
			if err != nil {
				//Error exists due to sending in same connection, figure it out
			}
			fmt.Print("Dotin server: ")
			log.Printf("Query mapping sent: %s\n", text)
			break
		}
	}
}
