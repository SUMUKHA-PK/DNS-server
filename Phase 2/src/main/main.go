package main

import (
	"fmt"
	"os"

	"../BytePacketBuffer"
	"../DNSPacket"
	"../errorHandling"
)

func main() {
	var buffer BytePacketBuffer.BytePacketBuffer

	f, err := os.Open("resp.txt")
	errorHandling.ErrorHandler(err)

	b1 := make([]byte, 512)
	n1, err := f.Read(b1)
	errorHandling.ErrorHandler(err)

	// fmt.Printf("%d bytes: %s\n", n1, string(b1))

	// _, err = f.Read(buffer.Buffer)
	// errorHandling.ErrorHandler(err)

	buffer.Buffer = b1[:n1]
	packet := DNSPacket.FromBuffer(buffer)

	fmt.Println(packet.Header)

	for i := 0; i < len(packet.Questions); i++ {
		fmt.Println(packet.Questions[i])
	}

	for i := 0; i < len(packet.Answers); i++ {
		fmt.Println(packet.Answers[i])
	}

	for i := 0; i < len(packet.Authorities); i++ {
		fmt.Println(packet.Authorities[i])
	}

	for i := 0; i < len(packet.Resources); i++ {
		fmt.Println(packet.Resources[i])
	}
}
