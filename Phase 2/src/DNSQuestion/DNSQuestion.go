package DNSQuestion

import (
	"../BytePacketBuffer"
	"../QueryType"
)

type DNSQuestion struct {
	name  string
	qtype QueryType.QueryType
}

func newQuestion(Question DNSQuestion, name string, QT QueryType.QueryType) DNSQuestion {

	Question.name = name
	Question.qtype = QT
	return Question
}

func read(Question DNSQuestion, Buffer BytePacketBuffer.BytePacketBuffer) {
	BytePacketBuffer.Read_qname(Buffer, Question.name)
	Question.qtype = QueryType.IntToQueryType(BytePacketBuffer.Read_u16(Buffer))
	// Class := BytePacketBuffer.Read_u16(Buffer) Dont know what this is
}
