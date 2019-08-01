package BytePacketBuffer

import (
	"errors"
	"fmt"
	"strings"

	"github.com/SUMUKHA-PK/DNS-server/Phase_2/errorHandling"
)

type BytePacketBuffer struct {
	Buffer []byte
	Pos    uint16
}

func New_BytePacktetBuffer(Buffer BytePacketBuffer) BytePacketBuffer {
	Buffer.Pos = 0
	return Buffer
}

func Pos(Buffer BytePacketBuffer) uint16 {
	return Buffer.Pos
}

func Step(Buffer BytePacketBuffer, steps uint16) error {
	if Buffer.Pos+steps >= 512 {
		return errors.New("End of buffer")
	}
	Buffer.Pos += steps
	return nil
}

func Seek(Buffer BytePacketBuffer, Pos uint16) error {
	if Pos >= 512 {
		return errors.New("End of buffer")
	}
	Buffer.Pos = Pos
	return nil
}

// Method reads one single byte and moves forward
func Read_data(Buffer BytePacketBuffer) (int, error) {
	if Buffer.Pos >= 512 {
		return 0, errors.New("End of buffer")
	}

	result := Buffer.Buffer[Buffer.Pos]
	Buffer.Pos += 1
	return int(result), nil
}

// Method for fetching data at a specified Position without modifying the internal Position

func Get(Buffer BytePacketBuffer, Pos uint16) (uint8, error) {
	if Pos >= 512 {
		return 0, errors.New("End of buffer")
	}

	result := Buffer.Buffer[Buffer.Pos]
	return uint8(result), nil
}

func Get_range(Buffer BytePacketBuffer, start uint16, len uint16) ([]byte, error) {
	if start+len >= 512 {
		return nil, errors.New("End of buffer")
	}
	return Buffer.Buffer[start : start+len], nil

}

func Read_u16(Buffer BytePacketBuffer) uint16 {
	a := uint16(errorHandling.ErrorHelper(Read_data(Buffer)))
	b := uint16(errorHandling.ErrorHelper(Read_data(Buffer)))
	result := (a << 8) | b
	return result

} // Error is raised by ErrorHelper

func Read_u32(Buffer BytePacketBuffer) uint32 {
	a := uint32(errorHandling.ErrorHelper(Read_data(Buffer)))
	b := uint32(errorHandling.ErrorHelper(Read_data(Buffer)))
	c := uint32(errorHandling.ErrorHelper(Read_data(Buffer)))
	d := uint32(errorHandling.ErrorHelper(Read_data(Buffer)))
	result := (a << 24) | (b << 16) | (c << 8) | (d << 0)
	return result
}

func Read_qname(Buffer BytePacketBuffer) string {

	var outstr string
	// Tracking current qname
	Pos := Buffer.Pos

	// to track a Possible jump
	jumped := false

	delimiter := ""

	for {
		// Start of a label has the length
		len, err := Get(Buffer, Pos)
		// fmt.Println(len)
		// fmt.Println(Pos)
		errorHandling.ErrorHandler(err)

		// If 2 most significant bits are set, indicates jump is needed
		if len&0xC0 == 0xC0 {
			// Jump to the place where the non-repeatitive data starts
			if !jumped {
				err := Seek(Buffer, Pos+2)
				errorHandling.ErrorHandler(err)
			}

			b_2, err := Get(Buffer, Pos+1)
			errorHandling.ErrorHandler(err)

			b2 := uint16(b_2)

			offset := uint16(((uint16(len) ^ 0xC0) << 8) | b2)
			fmt.Println(len)
			fmt.Println(b2)
			fmt.Println(offset)
			Pos = offset
			// Jump has happened
			jumped = true
		} else { //Base case, jump to Position and reading label, output values
			Pos += 1

			// Domain names are terminated by labels of length 0
			if len == 0 {
				break
			}

			// Append delimiter to output string
			outstr += delimiter

			str_buffer, err := Get_range(Buffer, Pos, uint16(len))
			errorHandling.ErrorHandler(err)

			outstr += strings.Trim(strings.Replace(fmt.Sprint(str_buffer), " ", "", -1), "[]")

			delimiter = "."

			//Move forward the entire length of the label
			Pos += uint16(len)
		}
	}
	// If jump has been performed weve already altered the buffer Position and musnt do it again
	if !jumped {
		err := Seek(Buffer, Pos)
		errorHandling.ErrorHandler(err)
	}

	return outstr
} // End of q_name
