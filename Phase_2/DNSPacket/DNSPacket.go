package DNSPacket

import (
	"fmt"

	"github.com/SUMUKHA-PK/DNS-server/Phase_2/BytePacketBuffer"
	"github.com/SUMUKHA-PK/DNS-server/Phase_2/DNSHeader"
	"github.com/SUMUKHA-PK/DNS-server/Phase_2/DNSQuestion"
	"github.com/SUMUKHA-PK/DNS-server/Phase_2/DNSRecord"
)

type DNSPacket struct {
	Header      DNSHeader.DNSHeader
	Questions   []DNSQuestion.DNSQuestion
	Answers     []DNSRecord.DNSRecord
	Authorities []DNSRecord.DNSRecord
	Resources   []DNSRecord.DNSRecord
}

func FromBuffer(Buffer BytePacketBuffer.BytePacketBuffer) DNSPacket {
	var result DNSPacket
	result.Header = DNSHeader.ReadHeader(result.Header, Buffer)

	fmt.Println(result.Header)
	for i := 0; i < int(result.Header.Questions); i++ {
		var question DNSQuestion.DNSQuestion
		question = DNSQuestion.Read(question, Buffer)
		result.Questions = append(result.Questions, question)
		fmt.Println(question)
	}

	for i := 0; i < int(result.Header.Answers); i++ {
		var record DNSRecord.DNSRecord
		record = DNSRecord.ReadRecord(Buffer, record)
		result.Answers = append(result.Answers, record)
		fmt.Println(record)
	}

	for i := 0; i < int(result.Header.Authoritative_entries); i++ {
		var record DNSRecord.DNSRecord
		record = DNSRecord.ReadRecord(Buffer, record)
		result.Authorities = append(result.Authorities, record)
	}

	for i := 0; i < int(result.Header.Resource_entries); i++ {
		var record DNSRecord.DNSRecord
		record = DNSRecord.ReadRecord(Buffer, record)
		result.Resources = append(result.Resources, record)
	}

	return result
}
