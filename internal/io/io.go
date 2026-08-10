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

func FileInDirectory(fileName string, path string) (bool, error) {
	files, err := os.ReadDir(path)
	if err != nil {
		return false, err
	}

	for _, file := range files {
		if file.Name() == fileName {
			return true, nil
		}
	}

	return false, nil
}
