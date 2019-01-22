package BytePacketBuffer

import (
	"errors"
	"fmt"
	"strings"

	"../errorHandling"
)

type BytePacketBuffer struct {
	buffer [512]int
	pos    int
}

func pos(Buffer BytePacketBuffer) int {
	return Buffer.pos
}

func step(Buffer BytePacketBuffer, steps int) error {
	if Buffer.pos+steps >= 512 {
		return errors.New("End of buffer")
	}
	Buffer.pos += steps
	return nil
}

func seek(Buffer BytePacketBuffer, pos int) error {
	if pos >= 512 {
		return errors.New("End of buffer")
	}
	Buffer.pos = pos
	return nil
}

// Method reads one single byte and moves forward
func read_data(Buffer BytePacketBuffer) (int, error) {
	if Buffer.pos >= 512 {
		return 0, errors.New("End of buffer")
	}

	result := Buffer.buffer[Buffer.pos]
	Buffer.pos += 1
	return result, nil
}

// Method for fetching data at a specified position without modifying the intenal position

func get(Buffer BytePacketBuffer, pos int) (uint8, error) {
	if pos >= 512 {
		return 0, errors.New("End of buffer")
	}

	result := Buffer.buffer[Buffer.pos]
	return uint8(result), nil
}

func get_range(Buffer BytePacketBuffer, start int, len int) ([]int, error) {
	if start+len >= 512 {
		return nil, errors.New("End of buffer")
	}
	return Buffer.buffer[start : start+len], nil

}

func read_u16(Buffer BytePacketBuffer) uint16 {
	a := uint16(errorHandling.ErrorHelper(read_data(Buffer)))
	b := uint16(errorHandling.ErrorHelper(read_data(Buffer)))
	result := (a << 8) | b
	return result

} // Error is raised by ErrorHelper

func read_u32(Buffer BytePacketBuffer) uint32 {
	a := uint32(errorHandling.ErrorHelper(read_data(Buffer)))
	b := uint32(errorHandling.ErrorHelper(read_data(Buffer)))
	c := uint32(errorHandling.ErrorHelper(read_data(Buffer)))
	d := uint32(errorHandling.ErrorHelper(read_data(Buffer)))
	result := (a << 24) | (b << 16) | (c << 8) | (d << 0)
	return result
}

func read_qname(Buffer BytePacketBuffer, outstr string) {

	// Tracking current qname
	pos := Buffer.pos

	// to track a possible jump
	jumped := false

	delimiter := ""

	for {
		// Start of a label has the length
		len, err := get(Buffer, pos)
		errorHandling.ErrorHandler(err)

		// If 2 most significant bits are set, indicates jump is needed
		if len&0xC0 == 0xC0 {
			// Jump to the place where the non-repeatitive data starts
			if !jumped {
				err := seek(Buffer, pos+2)
				errorHandling.ErrorHandler(err)
			}

			b_2, err := get(Buffer, pos+1)
			errorHandling.ErrorHandler(err)

			b2 := uint16(b_2)

			offset := ((uint16(len) ^ 0xC0) << 8) | b2

			pos = int(offset)
			// Jump has happened
			jumped = true
		} else { //Base case, jump to position and reading label, output values
			pos += 1

			// Domain names are terminated by labels of length 0
			if len == 0 {
				break
			}

			// Append delimiter to output string
			outstr += delimiter

			str_buffer, err := get_range(Buffer, pos, int(len))
			errorHandling.ErrorHandler(err)

			outstr += strings.Trim(strings.Replace(fmt.Sprint(str_buffer), " ", "", -1), "[]")

			delimiter = "."

			//Move forward the entire length of the label
			pos += int(len)
		}
	}
	// If jump has been performed weve already altered the buffer position and musnt do it again
	if !jumped {
		err := seek(Buffer, pos)
		errorHandling.ErrorHandler(err)
	}
}
