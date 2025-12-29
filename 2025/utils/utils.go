package utils

import (
	"bufio"
	"log"
	"os"
)

var PurpleFont string = ""

func ReadFile(filePath string) []string {
	fileName := filePath;

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}