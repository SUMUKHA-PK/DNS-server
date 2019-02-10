package DNSRecord

import (
	"../BytePacketBuffer"
	"../QueryType"
)

type DNSRecord struct {
}

func readRecord(Buffer BytePacketBuffer.BytePacketBuffer, Record DNSRecord) DNSRecord {

	domain := ""
	BytePacketBuffer.Read_qname(Buffer, domain)

	qtype_num := BytePacketBuffer.Read_u16(Buffer)
	qtype := QueryType.IntToQueryType(qtype_num)

	_ = BytePacketBuffer.Read_u16(Buffer) //Ignoring the class
	ttl := BytePacketBuffer.Read_u32(Buffer)
	data_len := BytePacketBuffer.Read_u16(Buffer)

	return Record
}
