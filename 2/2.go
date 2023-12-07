package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
)

func readFromFile() *os.File {
	file, err := os.Open("2.txt")
	if err != nil {
		log.Fatal(err)
	}
	return file
}

func split(r rune) bool{
	return r == ' ' || r == ',' || r == ':' || r == ';'
}


func main(){

	total := 0

	file := readFromFile()

	scanner := bufio.NewScanner(file)

	for scanner.Scan(){

		var cubeMap = map[string]int{
			"red": 0,
			"green": 0,
			"blue": 0,
		}

		text := scanner.Text()
		slc := strings.FieldsFunc(text, split)

		for i := 3; i < len(slc); i += 2 {

			value, _ := strconv.Atoi(slc[i-1])

			if value > cubeMap[slc[i]]{
				cubeMap[slc[i]] = value
			}
		}  

		subTotal := 1

		for _, value := range cubeMap {
			subTotal = subTotal * value
		}

		total += subTotal
		
	}

	fmt.Println(total)
	file.Close()
}