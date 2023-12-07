package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readFromFile() *os.File {
	fmt.Println("Test1")
	file, err := os.Open("1.txt")
	if err != nil {
		log.Fatal(err)
	}
	// defer file.Close()
	fmt.Println("Test12")
	return file
}

func main() {

	file := readFromFile()

	scanner := bufio.NewScanner(file)
	fmt.Println("Test3")
	for scanner.Scan() {
		fmt.Println("Test1")
		fmt.Println(scanner.Text())
	}
	file.Close()
}
