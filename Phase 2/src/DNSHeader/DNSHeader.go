package DNSHeader

import "../BytePacketBuffer"

//enumeration
type ResultCode uint8

const (
	FORMERR  ResultCode = 1
	SERVFAIL ResultCode = 2
	NXDOMAIN ResultCode = 3
	NOTIMP   ResultCode = 4
	REFUSED  ResultCode = 5
	NOERROR  ResultCode = 0
)

type DNSHeader struct {
	id uint16 //16bits

	recursion_desired    bool  // 1 bit
	truncated_message    bool  // 1 bit
	authoritative_answer bool  // 1 bit
	opcode               uint8 // 4 bits
	response             bool  // 1 bit

	rescode             ResultCode // 4 bits
	checking_disabled   bool       // 1 bit
	authed_data         bool       // 1 bit
	z                   bool       // 1 bit
	recursion_available bool       // 1 bit

	questions             uint16 // 16 bits
	answers               uint16 // 16 bits
	authoritative_entries uint16 // 16 bits
	resource_entries      uint16 // 16 bits
} // total 96 bits

func IntToResultCode(a uint8) ResultCode {
	if a == 0 {
		return NOERROR
	}
	if a == 1 {
		return FORMERR
	}
	if a == 2 {
		return SERVFAIL
	}
	if a == 3 {
		return NXDOMAIN
	}
	if a == 4 {
		return NOTIMP
	}
	return REFUSED
}

func new_Header(Header DNSHeader) DNSHeader {

	Header.id = 0

	Header.recursion_desired = false
	Header.truncated_message = false
	Header.authoritative_answer = false
	Header.opcode = 0
	Header.response = false

	Header.rescode = NOERROR
	Header.checking_disabled = false
	Header.authed_data = false
	Header.z = false
	Header.recursion_available = false

	Header.questions = 0
	Header.answers = 0
	Header.authoritative_entries = 0
	Header.resource_entries = 0

	return Header
}

func readHeader(Header DNSHeader, Buffer BytePacketBuffer.BytePacketBuffer) {

	Header.id = BytePacketBuffer.Read_u16(Buffer)

	flags := BytePacketBuffer.Read_u16(Buffer)
	a_a := flags >> 8
	a := uint8(a_a)
	b_b := flags & 0xFF
	b := uint8(b_b)
	Header.recursion_desired = (a & (1 << 0)) > 0
	Header.truncated_message = (a & (1 << 1)) > 0
	Header.authoritative_answer = (a & (1 << 0)) > 0
	Header.opcode = (a << 3) & 0x0F
	Header.response = (a & (1 << 7)) > 0

	Header.rescode = IntToResultCode(b & 0x0F)
	Header.checking_disabled = (b & (1 << 4)) > 0
	Header.authed_data = (b & (1 << 5)) > 0
	Header.z = (b & (1 << 6)) > 0
	Header.recursion_available = (b & (1 << 7)) > 0

	Header.questions = BytePacketBuffer.Read_u16(Buffer)
	Header.answers = BytePacketBuffer.Read_u16(Buffer)
	Header.authoritative_entries = BytePacketBuffer.Read_u16(Buffer)
	Header.resource_entries = BytePacketBuffer.Read_u16(Buffer)
}
