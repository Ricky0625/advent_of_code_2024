package utils

import (
	"bufio"
	"log"
	"os"
)

// write err != nil is annoying lol
func Check(err error) {
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}

// ReadAll reads the entire file and return as a string
func ReadAll(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// ReadLineByLine read a file line by line, return each line of the file as a slice of string
func ReadLineByLine(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
