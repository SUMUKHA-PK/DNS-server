package dotin

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

var receive string

func DotIn(IP_List_Name []string, IP_List_Addr []string) {

	//Starting the server
	link, err := net.Listen("tcp", "127.0.1.0:12345")
	if err != nil {
		fmt.Printf("Dotin server: ")
		fmt.Print(err)
	}

	//Continous server listening
	for {

		fmt.Printf("\nDotIn server listening for incoming connections on port 12345\n\n")
		conn, err := link.Accept()
		if err != nil {
			fmt.Printf("Dotin server: ")
			fmt.Print(err)
		}

		//Create a new scanner and get the data from the client
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			receive = scanner.Text()
			fmt.Printf("Dotin server: ")
			fmt.Printf("IP received to map from client: " + receive + "\n")
			break
		}
		if errReadConn := scanner.Err(); errReadConn != nil {
			fmt.Print(errReadConn)
			return
		}

		//Get the IP from the above servers
		result := get_data(receive, IP_List_Name, IP_List_Addr)

		//Communicate back the result to the client on the same connection
		scanner = bufio.NewScanner(strings.NewReader(result))

		for scanner.Scan() {
			text := scanner.Text()
			_, err := fmt.Fprintf(conn, text+"\n")
			if err != nil {
				//Error exists due to sending in same connection, figure it out
			}
			fmt.Printf("Dotin server: ")
			log.Print("Query mapping sent: " + text)
			break
		}
	}
}

func get_data(IP string, IP_List_Name []string, IP_List_Addr []string) string {
	split := strings.Split(IP, ".")

	var j, k int

	for i := 0; i < len(split); i++ {
		if split[i] == "in" {
			j = i
			break
		}
	}

	for i := 0; i < len(IP_List_Name); i++ {
		if split[j-1] == IP_List_Name[i] {
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
	fmt.Printf("Dotin server: ")
	fmt.Print("Client message: ")
	for scanner.Scan() {
		text := scanner.Text()
		_, errWrite := fmt.Fprintf(conn, text+"\n")
		if errWrite != nil {
			fmt.Print(err)
		}
		fmt.Printf("Dotin server: ")
		log.Print("IP sent to server: " + text)
		break
	}

	receive = get_data_next(IP, IP_List_Name, IP_List_Addr)

	// Receive mapping from the same connection
	scanner = bufio.NewScanner(conn)
	for scanner.Scan() {
		receive = scanner.Text()
		fmt.Printf("Dotin server: ")
		fmt.Printf("Mapping received: " + receive + "\n")
		break
	}
	if errReadConn := scanner.Err(); errReadConn != nil {
		fmt.Print(errReadConn)
	}

	return receive
}

func get_data_next(IP string, IP_List_Name []string, IP_List_Addr []string) string {
	split := strings.Split(IP, ".")

	var j, k int
	var str string

	for i := 0; i < len(split); i++ {
		if split[i] == "in" {
			j = i
			break
		}
	}

	str = "dot" + split[j-1]

	for i := 0; i < len(IP_List_Name); i++ {
		if str == IP_List_Name[i] {
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
	fmt.Printf("Dotedu server: ")
	fmt.Print("Client message: ")
	for scanner.Scan() {
		text := scanner.Text()
		_, errWrite := fmt.Fprintf(conn, text+"\n")
		if errWrite != nil {
			fmt.Print(err)
		}
		fmt.Printf("Dotedu server: ")
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
