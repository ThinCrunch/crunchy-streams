package io

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func GetFileLines(filepath string) []string {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var result []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}

	return result
}

func AskForString(prompt string) string {
	fmt.Printf(prompt)

	var result string
	_, err := fmt.Scan(&result)
	if err != nil {
		log.Fatal(err)
	}

	return result
}
