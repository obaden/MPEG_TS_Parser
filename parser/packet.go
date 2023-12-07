package parser

// Represents the header of a MPEG Transport Stream packet.
type MpegTsHeader struct {
	SyncByte byte
	Flags    byte
	PID      uint16
}
