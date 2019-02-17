package DNSPacket

import (
	"../BytePacketBuffer"
	"../DNSHeader"
	"../DNSQuestion"
	"../DNSRecord"
)

type DNSPacket struct {
	header      DNSHeader.DNSHeader
	questions   []DNSQuestion.DNSQuestion
	answers     []DNSRecord.DNSRecord
	authorities []DNSRecord.DNSRecord
	resources   []DNSRecord.DNSRecord
}

func FromBuffer(Buffer BytePacketBuffer.BytePacketBuffer) DNSPacket {
	var result DNSPacket
	var header DNSHeader.DNSHeader
	header = result.header
	x := header.questions
	DNSHeader.ReadHeader(header, Buffer)

	for i := 0; i < result.header.questions; i++ {

	}
}
