package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {

	fmt.Printf("Enter the query: ")
	var ip string
	fmt.Scan(&ip)

	conn, err := net.Dial("tcp", "192.168.43.10:12345")
	if err != nil {
		fmt.Print(err)
	} else {
		log.Print("Connected")
	}
	line := ip
	scanner := bufio.NewScanner(strings.NewReader(line))
	fmt.Print("Client message: ")
	for scanner.Scan() {
		text := scanner.Text()
		_, errWrite := fmt.Fprintf(conn, text+"\n")
		if errWrite != nil {
			fmt.Print(err)
		}
		log.Print("Server receives: " + text)
		break
	}
}
