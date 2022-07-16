package main

import (
	"bufio"
	"log"
	"os"
)

func ScanToken() string {
	file, err := os.Open("token")
	if err != nil {
		log.Fatalf("Unable to read token: %s", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		token = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return token
}