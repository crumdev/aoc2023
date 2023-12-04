package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	// retrieve list of strings from local file ./input.txt and split in an array by newline
	input, err := getInput("./input.txt")
	if err != nil {
		fmt.Printf("Error retrieving input: %s\n", err)
		return
	}

	// Iterate over the array and retrieve the first and last number from each string
	var sum int
	for _, line := range input {

		firstDigit, lastDigit, err := findFirstAndLastDigit(line)
		if err != nil {
			fmt.Printf("Error parsing digits: %s\n", err)
			return
		}
		concatenatedNumber := firstDigit + lastDigit
		concatenatedNumberInt, err := strconv.Atoi(concatenatedNumber)
		if err != nil {
			fmt.Printf("Error converting string to int: %s\n", err)
			return
		}
		sum += concatenatedNumberInt
	}

	fmt.Printf("Sum of retrieved numbers: %d\n", sum)
}

func findFirstAndLastDigit(s string) (string, string, error) {
	re := regexp.MustCompile(`\d`)
	matches := re.FindAllString(s, -1)

	if len(matches) == 0 {
		return "", "", fmt.Errorf("no digits found")
	}

	first := matches[0]
	last := matches[len(matches)-1]

	return first, last, nil
}

func getInput(filePath string) ([]string, error) {

	// input from local file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %s", err)
	}
	defer file.Close()

	// split file by newline into an array
	var input []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	return input, nil
}
