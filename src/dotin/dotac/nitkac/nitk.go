package nitkac

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

var receive string

func Nitkac(IP_List_Name []string, IP_List_Addr []string) {

	// Starting the server
	link, err := net.Listen("tcp", "127.0.3.0:12345")
	if err != nil {
		fmt.Printf("Nitkac server: ")
		fmt.Print(err)
	}

	//Continuos server listening
	for {
		fmt.Printf("\nNitkac Server listening for incoming connections on port 12345\n\n")
		conn, err := link.Accept()
		if err != nil {
			fmt.Printf("Nitkac server: ")
			fmt.Print(err)
		}

		//Create a new scanner and get the data from the client
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			receive = scanner.Text()
			fmt.Printf("Nitkac server: ")
			fmt.Printf("IP received to map from client: " + receive + "\n")
			break
		}
		if errReadConn := scanner.Err(); errReadConn != nil {
			fmt.Print(errReadConn)
			return
		}

		//Get the result of the mapping from the servers
		result := get_data(receive)

		//Communicate back the result to the client on the same connection
		scanner = bufio.NewScanner(strings.NewReader(result))

		for scanner.Scan() {
			text := scanner.Text()
			_, err := fmt.Fprintf(conn, text+"\n")
			if err != nil {
				//Error exists due to sending in same connection, figure it out
			}
			fmt.Printf("Nitkac server: ")
			log.Print("Query mapping sent: " + text)
			break
		}
	}
}

func get_data(IP string) string {

	split := strings.Split(IP, ".")

	if split[0] == "eee" && split[1] == "nitk" {
		return "10.3.0.14"
	}
	if split[0] == "cse" && split[1] == "nitk" {
		return "10.3.0.15"
	}
	if split[0] == "ece" && split[1] == "nitk" {
		return "10.3.0.16"
	}
	return "INVALID QUERY"
}
