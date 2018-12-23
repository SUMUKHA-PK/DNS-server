package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "192.168.1.13:12345")
	if err != nil {
		fmt.Print(err)
	} else {
		log.Print("Connected")
	}
	line := "wwww.google.com"
	var out []byte
	copy(out[:], line)
	conn.Write(out)
	scanner := bufio.NewScanner(os.Stdin)
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
