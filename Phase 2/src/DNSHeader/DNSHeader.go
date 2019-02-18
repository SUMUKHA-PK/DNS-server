package DNSHeader

import (
	"../BytePacketBuffer"
)

//enumeration
type ResultCode uint8

const (
	FORMERR  ResultCode = 1
	SERVFAIL ResultCode = 2
	NXDOMAIN ResultCode = 3
	NOTIMP   ResultCode = 4
	REFUSED  ResultCode = 5
	NOERROR  ResultCode = 0
) //enum

type DNSHeader struct {
	Id uint16 //16bits

	Recursion_desired    bool  // 1 bit
	Truncated_message    bool  // 1 bit
	Authoritative_answer bool  // 1 bit
	Opcode               uint8 // 4 bits
	Response             bool  // 1 bit

	Rescode             ResultCode // 4 bits
	Checking_disabled   bool       // 1 bit
	Authed_data         bool       // 1 bit
	Z                   bool       // 1 bit
	Recursion_available bool       // 1 bit

	Questions             uint16 // 16 bits
	Answers               uint16 // 16 bits
	Authoritative_entries uint16 // 16 bits
	Resource_entries      uint16 // 16 bits
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

func New_Header(Header DNSHeader) DNSHeader {

	Header.Id = 0

	Header.Recursion_desired = false
	Header.Truncated_message = false
	Header.Authoritative_answer = false
	Header.Opcode = 0
	Header.Response = false

	Header.Rescode = NOERROR
	Header.Checking_disabled = false
	Header.Authed_data = false
	Header.Z = false
	Header.Recursion_available = false

	Header.Questions = 0
	Header.Answers = 0
	Header.Authoritative_entries = 0
	Header.Resource_entries = 0

	return Header
}

func ReadHeader(Header DNSHeader, Buffer BytePacketBuffer.BytePacketBuffer) DNSHeader {

	Header.Id = BytePacketBuffer.Read_u16(Buffer)

	flags := BytePacketBuffer.Read_u16(Buffer)
	a_a := flags >> 8
	a := uint8(a_a)
	b_b := flags & 0xFF
	b := uint8(b_b)
	Header.Recursion_desired = (a & (1 << 0)) > 0
	Header.Truncated_message = (a & (1 << 1)) > 0
	Header.Authoritative_answer = (a & (1 << 0)) > 0
	Header.Opcode = (a << 3) & 0x0F
	Header.Response = (a & (1 << 7)) > 0

	Header.Rescode = IntToResultCode(b & 0x0F)
	Header.Checking_disabled = (b & (1 << 4)) > 0
	Header.Authed_data = (b & (1 << 5)) > 0
	Header.Z = (b & (1 << 6)) > 0
	Header.Recursion_available = (b & (1 << 7)) > 0

	Header.Questions = BytePacketBuffer.Read_u16(Buffer)
	Header.Answers = BytePacketBuffer.Read_u16(Buffer)
	Header.Authoritative_entries = BytePacketBuffer.Read_u16(Buffer)
	Header.Resource_entries = BytePacketBuffer.Read_u16(Buffer)

	return Header
}
