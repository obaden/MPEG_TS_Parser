package parser

import (
	"errors"
	"io"
)

// Reads the initial byte, discarding everything before the first sync byte
// to allow for reading a partial first packet.
//
// Returns the offset to the first byte, the buffer containing the first full packet and any errors
// if something went wrong reading the packet.
func readInitialPackage(input io.Reader) (int, [MpegPacketSize]byte, error) {
	var buffer [MpegPacketSize]byte

	for offset := 0; offset <= len(buffer); offset++ {
		_, err := input.Read(buffer[:1])
		if err != nil {
			return offset, buffer, err
		}

		// if we find the syncByte, read the remaining bytes of the package and return
		if buffer[0] == MpegSyncByte {
			_, err := io.ReadFull(input, buffer[1:])
			return offset, buffer, err
		}

		offset++
	}

	return len(buffer), buffer, errors.New("No sync byte present")
}

// Parses the header of a MPEG Transport Stream packet and return a header object.
func parseHeader(data [MpegPacketSize]byte) MpegTsHeader {
	var header MpegTsHeader

	// first byte in the packet is the sync byte
	header.SyncByte = data[0]

	// flags are contained in the first 3 bits of the second byte
	header.Flags = (data[1] & 0b11100000) >> 5

	// PID is made up of the last 5 bits of the second byte and all 8 bits of the third byte
	header.PID = 0b0001111111111111 & (uint16(data[1])<<8 | uint16(data[2]))

	return header
}

// Checks the packet header to make sure it represents a valid packet.
func isValidPacket(header MpegTsHeader) (bool, error) {
	if header.SyncByte != MpegSyncByte {
		return false, errors.New("No sync byte present")
	}

	return true, nil
}
