package root_server

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

var receive string

func Root_server(IP_List_Name []string, IP_List_Addr []string) {

	// Starting the server
	link, err := net.Listen("tcp", "127.0.0.1:12345")
	if err != nil {
		fmt.Printf("Root server: ")
		fmt.Print(err)
	}

	//Continous server listening
	for {
		fmt.Printf("\nRoot Server listening for incoming connections on port 12345\n\n")
		conn, err := link.Accept()
		if err != nil {
			fmt.Printf("Root server: ")
			fmt.Print(err)
		}

		//Create a new scanner and get the data from the client
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			receive = scanner.Text()
			fmt.Printf("Root server: ")
			fmt.Printf("IP received to map from client this: " + receive + "\n")
			break
		}
		if errReadConn := scanner.Err(); errReadConn != nil {
			fmt.Print(errReadConn)
			fmt.Printf("Root server: ")
			return
		}

		//Get the result of the mapping from the servers
		result := start_servers(receive, IP_List_Name, IP_List_Addr)

		//Communicate back the result to the client on the same connection
		scanner = bufio.NewScanner(strings.NewReader(result))

		for scanner.Scan() {
			text := scanner.Text()
			_, err := fmt.Fprintf(conn, text+"\n")
			if err != nil {
				//Error exists due to sending in same connection, figure it out
			}
			fmt.Printf("Root server: ")
			log.Print("Query mapping sent: " + text)
			break
		}
	}
}

func start_servers(IP string, IP_List_Name []string, IP_List_Addr []string) string {
	log.Printf("Starting servers, sending data: \n")

	split := strings.Split(IP, ".")
	fmt.Printf("Length of split: %d\n", len(split))

	var k int
	var str string

	if split[len(split)-1] == "com" {
		str = "dotcom"
	} else if split[len(split)-1] == "edu" {
		str = "dotedu"
	} else if split[len(split)-1] == "in" {
		str = "dotin"
	}

	fmt.Print(str)

	for i := 0; i < len(IP_List_Name); i++ {
		if IP_List_Name[i] == str {
			k = i
			break
		}
	}

	addr := IP_List_Addr[k]

	addr += ":12345"
	log.Printf(addr)

	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Print(err)
	} else {
		log.Print("Connected")
	}

	// Send the query to the root server
	line := IP
	scanner := bufio.NewScanner(strings.NewReader(line))
	fmt.Print("Client message: ")
	for scanner.Scan() {
		text := scanner.Text()
		_, errWrite := fmt.Fprintf(conn, text+"\n")
		if errWrite != nil {
			fmt.Print(err)
		}
		log.Print("IP sent to server: " + text)
		break
	}

	// Receive mapping from the same connection
	scanner = bufio.NewScanner(conn)
	for scanner.Scan() {
		receive = scanner.Text()
		fmt.Printf("Mapping received: " + receive + "\n")

		break
	}
	if errReadConn := scanner.Err(); errReadConn != nil {
		fmt.Print(errReadConn)
	}

	return receive
}
