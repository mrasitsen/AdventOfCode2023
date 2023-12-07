package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

var numberMap = map[string]int{
	"one": 1,
	"two": 2,
	"three": 3,
	"four": 4,
	"five": 5,
	"six": 6,
	"seven": 7,
	"eight": 8,
	"nine": 9,
}

func readFromFile() *os.File {
	file, err := os.Open("1.txt")
	if err != nil {
		log.Fatal(err)
	}
	return file
}

func containsSubstring(text string) (int, bool) {

	for key, value := range numberMap {
		if len(text) >= len(key) {
			if text[0:len(key)] == key {
				return value, true
			}
		}
	}

	return 0, false
}

func main() {

	file := readFromFile()

	var total int = 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		var firstDigit int = 0
		var lastDigit int = 0

		var number int = 0
		isFirst := true

		text := scanner.Text()

		for index, char := range text {
			if unicode.IsDigit(char) {
				
				number = int(char - '0')
			}
			
			value, result := containsSubstring(text[index:len(text)-1])
			if result {
				number = value
			}

			if (number != 0) {
				if isFirst {
					firstDigit = number
					isFirst = false
				}
				lastDigit = number
			}
			
		}
		total = total + ((firstDigit * 10) + lastDigit)
	}

	fmt.Println(total)

	file.Close()
}
