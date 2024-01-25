package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"unicode"
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

	err, output = translate(&input, ituMap)
	if err != nil {
		log.Fatal(err)
	}
	println(*output, lSpace)
}

func translate(input *string, charMap map[string]string) (error, *string) {
	var err error
	var output string
	for _, c := range *input {
		normalizedC := string(unicode.ToLower(c))
		val, ok := charMap[normalizedC]
		if !ok {
			err = errors.New("Unsupported character: " + normalizedC)
		}
		output += val
		output += "\u0020"
	}

	return err, &output
}
