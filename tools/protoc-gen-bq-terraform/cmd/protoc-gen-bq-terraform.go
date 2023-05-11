package main

import (
	"flag"
	"log"
)

func main() {
	// Parse all flags
	flag.Parse()

	// Load the configuration
	bqConfiguration := NewBQConfiguration()

	// Print debug info if enabled
	bqConfiguration.PrintDebug()

	// If valid, write to file
	if bqConfiguration.IsValid() {
		bqConfiguration.WriteFile()
	} else {
		log.Fatal("Configuration is not valid")
	}
}