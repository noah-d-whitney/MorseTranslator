package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Welcome to morse translator")
	fmt.Println("Enter english text")

	var input string
	var output *string

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	input = scanner.Text()

	err, output = translateEToItu(&input)
	if err != nil {
		log.Fatal(err)
	}
	println(*output, lSpace)
}
