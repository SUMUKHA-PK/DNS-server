package dell

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

var receive string

func Dell(IP_List_Name []string, IP_List_Addr []string) {

	// Starting the server
	link, err := net.Listen("tcp", "127.0.2.3:12345")
	if err != nil {
		fmt.Printf("Dell server: ")
		fmt.Print(err)
	}

	//Continuos server listening
	for {

		fmt.Printf("\nDell Server listening for incoming connections on port 12345\n\n")
		conn, err := link.Accept()
		if err != nil {
			fmt.Print("Dell server: ")
			fmt.Print(err)
		}

		//Create a new scanner and get the data from the client
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			receive = scanner.Text()
			fmt.Print("Dell server: ")
			fmt.Printf("IP received to map from client: %s\n", receive)
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
			fmt.Print("Dell server: ")
			log.Printf("Query mapping sent: %s\n", text)
			break
		}
	}
}

func get_data(IP string) string {

	split := strings.Split(IP, ".")

	if split[0] == "www" && split[1] == "dell" {
		return "143.166.147.101"
	}
	return "INVALID QUERY"
}
