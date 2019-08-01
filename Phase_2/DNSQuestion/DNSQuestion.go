package DNSQuestion

import (
	"github.com/SUMUKHA-PK/DNS-server/Phase_2/BytePacketBuffer"
	"github.com/SUMUKHA-PK/DNS-server/Phase_2/QueryType"
)

type DNSQuestion struct {
	Name  string
	Qtype QueryType.QueryType
}

func NewQuestion(Question DNSQuestion, name string, QT QueryType.QueryType) DNSQuestion {
	Question.Name = name
	Question.Qtype = QT
	return Question
}

func Read(Question DNSQuestion, Buffer BytePacketBuffer.BytePacketBuffer) DNSQuestion {
	Question.Name = BytePacketBuffer.Read_qname(Buffer)
	Question.Qtype = QueryType.IntToQueryType(BytePacketBuffer.Read_u16(Buffer))
	_ = BytePacketBuffer.Read_u16(Buffer) //Dont know what this is
	return Question
}
