package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

func getFirstAndLastDigitOfString(inputString string) (int, int) {
	firstDigitFound := false
	firstDigit := 0
	lastDigit := 0

	for _, char := range inputString {
		if !unicode.IsDigit(char) {
			continue
		}

		if !firstDigitFound {
			firstDigit = int(char - '0')
			lastDigit = firstDigit //the first and last digit can overlap so, we assign the firstDigit to the last
			firstDigitFound = true
			continue
		}

		lastDigit = int(char - '0')
	}

	return firstDigit, lastDigit
}

func main() {
	file, err := os.Open("./trebuchet_data.txt")
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

	fmt.Println(outputDigitsFound)

	var outputSums []int
	for _, v := range outputDigitsFound {
		outputSums = append(outputSums, (v[0]*10)+v[1])
	}
	fmt.Println(outputSums)

	totalSum := 0

	for _, v := range outputSums {
		totalSum = v + totalSum
	}

	fmt.Println(totalSum)
}
