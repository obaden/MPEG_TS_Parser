package parser

import (
	"fmt"
	"io"
	"sort"
)

func ParseStream(input io.Reader) ([]uint16, error) {
	uniquePIDs := make(map[uint16]struct{})

	var buffer [MpegPacketSize]byte
	var offset int
	var packetNumber int = 0

	offset, buffer, err := readInitialPackage(input)
	if err != nil && err != io.ErrUnexpectedEOF {
		return nil, formatError(err, packetNumber, offset)
	}

	for {

		// Parse the packet
		header := parseHeader(buffer)

		// Validate packet
		packetValid, err := isValidPacket(header)
		if !packetValid {
			return nil, formatError(err, packetNumber, offset)
		}

		// Store unique PIDs
		if _, exists := uniquePIDs[header.PID]; !exists {
			uniquePIDs[header.PID] = struct{}{}
		}

		// Read the next packet
		bytesRead, err := io.ReadFull(input, buffer[:])
		if err == io.EOF || err == io.ErrUnexpectedEOF {
			break
		} else if err != nil {
			return nil, err
		}

		offset += bytesRead
		packetNumber++
	}

	return convertToSortedSlice(uniquePIDs), nil
}

func formatError(err error, packetCount int, offset int) error {
	return fmt.Errorf("%s in packet %d, offset %d", err, packetCount, offset)
}

func convertToSortedSlice(uniquePIDs map[uint16]struct{}) []uint16 {
	pids := make([]uint16, 0, len(uniquePIDs))
	for pid := range uniquePIDs {
		pids = append(pids, pid)
	}

	sort.Slice(pids, func(i, j int) bool { return pids[i] < pids[j] })

	return pids
}
