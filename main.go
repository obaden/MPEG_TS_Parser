package main

import (
	"fmt"
	"os"

	"github.com/obaden/MPEG_TS_Parser/parser"
)

func main() {
	uniquePIDs, err := parser.ParseStream(os.Stdin)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// Print unique PIDs
	for _, pid := range uniquePIDs {
		fmt.Printf("0x%04X\n", pid)
	}

	os.Exit(0)
}
