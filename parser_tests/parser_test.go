package parser_tests

import (
	"bytes"
	"encoding/hex"
	"testing"

	"github.com/obaden/MPEG_TS_Parser/parser"
)

func TestPartialPackageWithCorrectSyncByte(t *testing.T) {
	input, _ := hex.DecodeString("47000101010101")
	inputStream := bytes.NewReader(input)

	pids, err := parser.ParseStream(inputStream)

	if err != nil {
		t.Errorf("ParseStream returned error: %v", err)
	}

	// Replace with your expected pids
	expectedPids := []uint16{0x01}

	if !equal(pids, expectedPids) {
		t.Errorf("ParseStream returned %v, want %v", pids, expectedPids)
	}
}

func TestPartialPackageAtStartOfStream(t *testing.T) {
	input, _ := hex.DecodeString("0101010147000101010101")
	inputStream := bytes.NewReader(input)

	pids, err := parser.ParseStream(inputStream)

	if err != nil {
		t.Errorf("ParseStream returned error: %v", err)
	}

	// Replace with your expected pids
	expectedPids := []uint16{0x01}

	if !equal(pids, expectedPids) {
		t.Errorf("ParseStream returned %v, want %v", pids, expectedPids)
	}
}

func TestMissingSyncByte(t *testing.T) {
	input, _ := hex.DecodeString("01010101010101")
	inputStream := bytes.NewReader(input)

	_, err := parser.ParseStream(inputStream)

	if err == nil {
		t.Errorf("ParseStream did not return error. Expected No Sync Byte present error")
	}
}

// Helper function to compare slices
func equal(a, b []uint16) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
