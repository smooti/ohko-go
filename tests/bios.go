package main

import (
	"fmt"
	"log"

	"github.com/digitalocean/go-smbios/smbios"
)

func main() {
	rc, _, err := smbios.Stream()
	if err != nil {
		log.Fatalf("error getting SMBIOS stream: %v", err)
	}
	defer rc.Close()

	decoder := smbios.NewDecoder(rc)
	// fmt.Println(decoder)
	entries, err := decoder.Decode()
	fmt.Println(entries)
	if err != nil {
		log.Fatalf("error decoding SMBIOS structures: %v", err)
	}

	for i, entry := range entries {
		fmt.Printf("Entry %d: | Header type: %d | Strings: %s\n", i, entry.Header.Type, entry.Strings)
	}
}
