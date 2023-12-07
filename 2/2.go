package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
)

var cubeMap = map[string]int{
	"red": 12,
	"green": 13,
	"blue": 14,
}

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

		isOk := true
		text := scanner.Text()
		slc := strings.FieldsFunc(text, split)
		gameNumb := slc[1]

		for i := 3; i < len(slc); i += 2 {

			value, _ := strconv.Atoi(slc[i-1])

			if value > cubeMap[slc[i]]{
				isOk = false
				break
			}
		}  

		if isOk {
			num, _ := strconv.Atoi(gameNumb)
			total = total + num
		}
		
	}

	fmt.Println(total)
	file.Close()
}