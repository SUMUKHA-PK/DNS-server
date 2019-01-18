package unit

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func Unit() {
	var IP_Name = []string{"cse.nitk.ac.in", "ece.nitk.ac.in", "eee.nitk.ac.in", "www.google.in", "www.lynda.com", "www.amazon.net", "www.dell.in", "www.amazon.com", "fsae.nitk.ac.in", "www.dell.com", "www.nitk.edu", "fre.nitk.edu", "www.mit.edu"}
	var IP_Addr = []string{"10.3.0.15", "10.3.0.16", "10.3.0.14", "216.58.207.163", "INVALID QUERY", "INVALID QUERY", "INVALID QUERY", "52.85.180.219", "INVALID QUERY", "143.166.147.101", "10.100.12.123", "INVALID QUERY", "23.76.235.103"}

	var count, total int

	count = 0
	total = 0

	for i := 0; i < len(IP_Addr); i++ {

		var receive string

		ip := IP_Name[i]
		//Root server is always hosted at this
		conn, err := net.Dial("tcp", "127.0.0.1:12345")
		if err != nil {
			fmt.Print(err)
		} else {
			log.Print("Connected")
		}

		// Send the query to the root server
		line := ip
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
			return
		}
		if receive == IP_Addr[i] {
			log.Printf("Test passed!")
			count++
		} else {
			log.Printf("Test failed!")
		}
		total++
	}
	fmt.Printf("\n%d off %d tests passed\n", count, total)
}
