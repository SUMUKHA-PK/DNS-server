package dell

import (
	"fmt"
)

var receive string

func Dell() {
	fmt.Printf("I am the dell authoritative server!\n")

	// // Starting the server
	// link, err := net.Listen("tcp", "127.0.2.3:12345")
	// if err != nil {
	// 	fmt.Print(err)
	// 	fmt.Printf("\n2\n")
	// }

	// for {
	// 	fmt.Printf("\nRoot Server listening for incoming connections on port 12345\n\n")
	// 	conn, err := link.Accept()
	// 	if err != nil {
	// 		fmt.Print(err)
	// 		fmt.Printf("\n1\n")
	// 	}

	// 	//Create a new scanner and get the data from the client
	// 	scanner := bufio.NewScanner(conn)
	// 	for scanner.Scan() {
	// 		receive = scanner.Text()
	// 		fmt.Printf("IP received to map from client: " + receive + "\n")

	// 		break
	// 	}
	// 	if errReadConn := scanner.Err(); errReadConn != nil {
	// 		fmt.Print(errReadConn)
	// 		return
	// 	}

	// 	//Get the result of the mapping from the servers
	// 	result := start_servers(receive)

	// 	//Communicate back the result to the client on the same connection
	// 	scanner = bufio.NewScanner(strings.NewReader(result))

	// 	for scanner.Scan() {
	// 		text := scanner.Text()
	// 		_, err := fmt.Fprintf(conn, text+"\n")
	// 		if err != nil {
	// 			//Error exists due to sending in same connection, figure it out
	// 		}
	// 		log.Print("Query mapping sent: " + text)
	// 		break
	// 	}
	// }
}
