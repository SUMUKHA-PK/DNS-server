package helper

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

var receive string

func ContactHelper(IP string, IP_List_Name []string, IP_List_Addr []string, server_name string) string {
	split := strings.Split(IP, ".")

	var j int

	//Error control mechanism
	k := -1

	var str, str1 string

	for i := 0; i < len(split); i++ {
		if split[i] == server_name {
			j = i
			break
		}
	}

	str = "dot" + split[j-1]
	str1 = split[j-1]

	for i := 0; i < len(IP_List_Name); i++ {
		if str == IP_List_Name[i] || str1 == IP_List_Name[i] {
			k = i
			break
		}
	}

	if k == -1 {
		log.Printf("No valid IP at %s server", server_name)

		// defer func() {
		// 	if r := recover(); r != nil {
		// 		log.Fatal(r)
		// 	}
		// }()

		// panic("INVALID QUERY")
		return "INVALID QUERY"
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

	// Send the query to the dotac server
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
		fmt.Printf("%s server: ", server_name)
		log.Print("IP sent to server: " + text)
		break
	}

	// Receive mapping from the same connection
	scanner = bufio.NewScanner(conn)
	for scanner.Scan() {
		receive = scanner.Text()
		fmt.Printf("%s server: ", server_name)
		fmt.Printf("Mapping received: " + receive + "\n")
		break
	}
	if errReadConn := scanner.Err(); errReadConn != nil {
		fmt.Print(errReadConn)
	}

	return receive
}
