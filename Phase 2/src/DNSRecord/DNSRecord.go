package DNSRecord

import (
	"net"

	"../BytePacketBuffer"
	"../QueryType"
	"../errorHandling"
)

type DNSRecord struct {
	domain   string
	addr     net.IP
	ttl      uint32
	qtype    uint16
	data_len uint16
}

func readRecord(Buffer BytePacketBuffer.BytePacketBuffer, Record DNSRecord) DNSRecord {

	domain := ""
	BytePacketBuffer.Read_qname(Buffer, domain)

	qtype_num := BytePacketBuffer.Read_u16(Buffer)
	qtype := QueryType.IntToQueryType(qtype_num)

	_ = BytePacketBuffer.Read_u16(Buffer) //Ignoring the class
	ttl := BytePacketBuffer.Read_u32(Buffer)
	data_len := BytePacketBuffer.Read_u16(Buffer)

	if qtype == QueryType.A {
		raw_addr := BytePacketBuffer.Read_u32(Buffer)
		addr := net.IPv4(uint8((raw_addr>>24)&0xFF),
			uint8((raw_addr>>16)&0xFF),
			uint8((raw_addr>>8)&0xFF),
			uint8((raw_addr>>0)&0xFF))

		Record.domain = domain
		Record.addr = addr
		Record.ttl = ttl
	} else {
		err := BytePacketBuffer.Step(Buffer, data_len)
		errorHandling.ErrorHandler(err)
		Record.domain = domain
		Record.qtype = qtype_num
		Record.data_len = data_len
		Record.ttl = ttl
	}
	return Record
}
