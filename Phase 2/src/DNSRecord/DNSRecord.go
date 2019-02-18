package DNSRecord

import (
	"fmt"
	"net"

	"../BytePacketBuffer"
	"../QueryType"
	"../errorHandling"
)

type DNSRecord struct {
	Domain   string
	Addr     net.IP
	Ttl      uint32
	Qtype    uint16
	Data_len uint16
}

func ReadRecord(Buffer BytePacketBuffer.BytePacketBuffer, Record DNSRecord) DNSRecord {

	domain := BytePacketBuffer.Read_qname(Buffer)

	qtype_num := BytePacketBuffer.Read_u16(Buffer)
	qtype := QueryType.IntToQueryType(qtype_num)

	_ = BytePacketBuffer.Read_u16(Buffer) //Ignoring the class
	ttl := BytePacketBuffer.Read_u32(Buffer)
	data_len := BytePacketBuffer.Read_u16(Buffer)

	fmt.Print("1")
	if qtype == QueryType.A {
		raw_addr := BytePacketBuffer.Read_u32(Buffer)
		addr := net.IPv4(uint8((raw_addr>>24)&0xFF),
			uint8((raw_addr>>16)&0xFF),
			uint8((raw_addr>>8)&0xFF),
			uint8((raw_addr>>0)&0xFF))

		Record.Domain = domain
		Record.Addr = addr
		Record.Ttl = ttl
		fmt.Print("12")
	} else {
		err := BytePacketBuffer.Step(Buffer, data_len)
		errorHandling.ErrorHandler(err)
		Record.Domain = domain
		Record.Qtype = qtype_num
		Record.Data_len = data_len
		Record.Ttl = ttl
		fmt.Print("2")
	}
	return Record
}
