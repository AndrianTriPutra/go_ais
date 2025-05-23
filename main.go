package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/BertoldVdb/go-ais"
	"github.com/BertoldVdb/go-ais/aisnmea"
)

func main() {
	nm := aisnmea.NMEACodecNew(ais.CodecNew(false, false))
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("line: %s\n", line)

		decoded, err := nm.ParseSentence(line)
		if err != nil {
			log.Fatalf("Failed parse: %v", err)
		}

		if decoded == nil || decoded.Packet == nil {
			continue
		}

		payload, err := json.MarshalIndent(decoded.Packet, "", "  ")
		if err != nil {
			log.Fatalf("Failed encode to JSON: %v", err)
		}
		fmt.Printf("payload: %s\n", payload)
	}
}
