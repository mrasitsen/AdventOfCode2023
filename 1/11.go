package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

func readFromFile() *os.File {
	file, err := os.Open("1.txt")
	if err != nil {
		log.Fatal(err)
	}
	return file
}

func main() {

	file := readFromFile()

	var total int = 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		var firstDigit int = 0
		var lastDigit int = 0
		isFirst := true

		text := scanner.Text()
		for _, char := range text {
			if unicode.IsDigit(char) {
				if isFirst {
					firstDigit = int(char - '0')
					isFirst = false
				}
				lastDigit = int(char - '0')
			}

		}
		total = total + ((firstDigit * 10) + lastDigit)
	}

	fmt.Println(total)

	file.Close()
}
