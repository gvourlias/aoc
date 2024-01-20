package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

// one, two, three, four, five, six, seven, eight, nine
var stringNumbersToNumbers = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

var stringNumbersLength = [3]int{3, 4, 5}

func getFirstAndLastDigitOfString(inputString string) (int, int) {
	firstDigitFound := false
	firstDigit := 0
	lastDigit := 0

	for v, char := range inputString {
		var digit = 0
		if !unicode.IsDigit(char) {
			lengthOfString := len(inputString)
			lengthOfStringRemaining := len(inputString[v:lengthOfString])

			//it means that there isnt enough chars to make a number so we end this early
			if lengthOfStringRemaining < 3 {
				continue
			}

			found := false
			for i := 0; i < 3; i++ {
				upperBound := v + stringNumbersLength[i]
				if upperBound > lengthOfString {
					continue
				}

				fragment := inputString[v:upperBound]
				digitValue, ok := stringNumbersToNumbers[fragment]
				if !ok {
					continue
				}

				digit = digitValue
				found = true
				break
			}

			if !found {
				continue
			}
		} else {
			digit = int(char - '0')
		}

		if !firstDigitFound {
			firstDigit = digit
			lastDigit = firstDigit //the first and last digit can overlap so, we assign the firstDigit to the last
			firstDigitFound = true
			continue
		}
		lastDigit = digit
	}

	return firstDigit, lastDigit
}

func trebuchet(filePath string) int {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var outputDigitsFound [][]int
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) == 0 {
			continue
		}

		first, last := getFirstAndLastDigitOfString(text)
		outputDigitsFound = append(outputDigitsFound, []int{first, last})
	}

	// fmt.Println("First and last digits per line:", outputDigitsFound)

	var outputSums []int
	for _, v := range outputDigitsFound {
		outputSums = append(outputSums, (v[0]*10)+v[1])
	}
	// fmt.Println("First and last digits to numbers, and their respective sums:", outputSums)

	totalSum := 0

	for _, v := range outputSums {
		totalSum = v + totalSum
	}

	return totalSum
}

func main() {
	fmt.Println("Total sum:", trebuchet("./trebuchet_data.txt"))
}
