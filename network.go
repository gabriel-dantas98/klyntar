package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/Ullaakut/nmap"
)

func getOpenPorts() []string {
	var (
		resultBytes []byte
		errorBytes  []byte
	)
	// Equivalent to `/usr/local/bin/nmap -p 80,443,843 google.com facebook.com youtube.com`,
	// with a 5 minute timeout.
	s, err := nmap.NewScanner(
		nmap.WithTargets("localhost"),
		nmap.WithPorts("0-49151"),
	)
	if err != nil {
		log.Fatalf("unable to create nmap scanner: %v", err)
	}

	// Executes asynchronously, allowing results to be streamed in real time.
	if err := s.RunAsync(); err != nil {
		panic(err)
	}

	// Connect to stdout of scanner.
	stdout := s.GetStdout()

	// Connect to stderr of scanner.
	stderr := s.GetStderr()

	// Goroutine to watch for stdout and print to screen. Additionally it stores
	// the bytes intoa variable for processiing later.
	go func() {
		for stdout.Scan() {
			// fmt.Println(stdout.Text())
			resultBytes = append(resultBytes, stdout.Bytes()...)
		}
	}()

	// Goroutine to watch for stderr and print to screen. Additionally it stores
	// the bytes intoa variable for processiing later.
	go func() {
		for stderr.Scan() {
			errorBytes = append(errorBytes, stderr.Bytes()...)
		}
	}()

	// Blocks main until the scan has completed.
	if err := s.Wait(); err != nil {
		panic(err)
	}

	// Parsing the results into corresponding structs.
	result, err := nmap.Parse(resultBytes)

	// Parsing the results into the NmapError slice of our nmap Struct.
	result.NmapErrors = strings.Split(string(errorBytes), "\n")
	if err != nil {
		panic(err)
	}

	var openPorts []string

	// Use the results to print an example output
	for _, host := range result.Hosts {
		if len(host.Ports) == 0 || len(host.Addresses) == 0 {
			continue
		}

		for _, port := range host.Ports {
			if fmt.Sprintf("%s", port.State) == "open" {
				// fmt.Printf("\tPort %d/%s %s %s\n", port.ID, port.Protocol, port.State, port.Service.Name)
				openPorts = append(openPorts, fmt.Sprintf("%d", port.ID))
			}
		}
	}

	return openPorts
}
