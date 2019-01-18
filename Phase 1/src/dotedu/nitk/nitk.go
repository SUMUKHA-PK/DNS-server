package nitk

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

var receive string

func Nitk(IP_List_Name []string, IP_List_Addr []string) {

	// Starting the server
	link, err := net.Listen("tcp", "127.0.2.4:12345")
	if err != nil {
		fmt.Printf("Nitk server: ")
		fmt.Print(err)
	}

	//Continuous server listening
	for {

		fmt.Printf("\nNITK edu Server listening for incoming connections on port 12345\n\n")
		conn, err := link.Accept()
		if err != nil {
			fmt.Printf("Nitk server: ")
			fmt.Print(err)
		}

		//Create a new scanner and get the data from the client
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			receive = scanner.Text()
			fmt.Printf("Nitk server: ")
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
			fmt.Printf("Nitk server: ")
			log.Print("Query mapping sent: " + text)
			break
		}
	}
}

func get_data(IP string) string {

	split := strings.Split(IP, ".")

	if split[0] == "www" && split[1] == "nitk" {
		return "10.100.12.123"
	}
	return "INVALID QUERY"
}
