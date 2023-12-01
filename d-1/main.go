package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func extractNumber(s string) ([]int, error) {
	re := regexp.MustCompile(`\d`)
	matches := re.FindAllString(s, -1)

	if len(matches) == 0 {
		log.Println("number not found on this string")
		return nil, fmt.Errorf("number not found on this string")
	}

	var numbers []int

	if len(matches) == 1 {
		number, err := strconv.Atoi(matches[0])
		if err != nil {
			log.Println("failed to convert number [1]")
			return nil, err
		}
		numbers = append(numbers, number, number)
	} else {
		for _, match := range matches {
			number, err := strconv.Atoi(match)
			if err != nil {
				log.Println("failed to convert number [2]")
				return nil, err
			}
			numbers = append(numbers, number)
		}
	}

	return numbers, nil
}

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer input.Close()

	scanner := bufio.NewScanner(input)

	var totalNumber int
	for scanner.Scan() {
		numArray, err := extractNumber(scanner.Text())
		log.Printf("numArray = %+v", numArray)
		if err != nil {
			totalNumber += 0
		}
		firstVal := numArray[0]
		lastVal := numArray[len(numArray)-1]

		combinedStr := strconv.Itoa(firstVal) + strconv.Itoa(lastVal)
		combinedNumber, err := strconv.Atoi(combinedStr)
		if err != nil {
			log.Fatal(err)
		}

		totalNumber += combinedNumber
	}

	fmt.Printf("Total Number = %d", totalNumber)
}
