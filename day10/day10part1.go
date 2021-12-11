package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	chunks := parseInput("input.txt")

	badCharacterScore := 0
	for _, chunk := range chunks {

		closingCharacters := make([]string, 0)
		valid := true
		// fmt.Printf("Chuck: %v\n", chunk)
		characters := strings.Split(chunk, "")
		for i := 0; i < len(characters); i++ {
			if valid == false {
				break
			}
			switch characters[i] {
			case "(":
				closingCharacters = append(closingCharacters, "(")
				// fmt.Printf("Add Characters    %v\n", closingCharacters)
			case "[":
				closingCharacters = append(closingCharacters, "[")
				// fmt.Printf("Add Characters    %v\n", closingCharacters)
			case "{":
				closingCharacters = append(closingCharacters, "{")
				// fmt.Printf("Add Characters    %v\n", closingCharacters)
			case "<":
				closingCharacters = append(closingCharacters, "<")
				// fmt.Printf("Add Characters    %v\n", closingCharacters)
			case ")":
				if closingCharacters[len(closingCharacters)-1] == "(" {
					// fmt.Printf("Delete Characters %v\n", closingCharacters[len(closingCharacters)-1])
					closingCharacters = RemoveIndex(closingCharacters, len(closingCharacters)-1)
				} else {
					valid = false
					badCharacterScore += 3
					fmt.Printf("Bad Character Found: %s\n", characters[i])
					break
				}
			case "]":
				if closingCharacters[len(closingCharacters)-1] == "[" {
					// fmt.Printf("Delete Characters %v\n", closingCharacters[len(closingCharacters)-1])
					// fmt.Printf("Delete Characters %v\n", closingCharacters)
					closingCharacters = RemoveIndex(closingCharacters, len(closingCharacters)-1)
				} else {
					valid = false
					badCharacterScore += 57
					fmt.Printf("Bad Character Found: %s\n", characters[i])
					break
				}
			case "}":
				if closingCharacters[len(closingCharacters)-1] == "{" {
					// fmt.Printf("Delete Characters %v\n", closingCharacters[len(closingCharacters)-1])
					// fmt.Printf("Delete Characters %v\n", closingCharacters)
					closingCharacters = RemoveIndex(closingCharacters, len(closingCharacters)-1)
				} else {
					valid = false
					badCharacterScore += 1197
					fmt.Printf("Bad Character Found: %s\n", characters[i])
					break
				}
			case ">":
				if closingCharacters[len(closingCharacters)-1] == "<" {
					// fmt.Printf("Delete Characters %v\n", closingCharacters[len(closingCharacters)-1])
					// fmt.Printf("Delete Characters %v\n", closingCharacters)
					closingCharacters = RemoveIndex(closingCharacters, len(closingCharacters)-1)
				} else {
					valid = false
					badCharacterScore += 25137
					fmt.Printf("Bad Character Found: %s\n", characters[i])
					break
				}
			}
		}
		fmt.Printf("String %s is %v\n", chunk, valid)
	}
	fmt.Printf("Syntax Score: %d", badCharacterScore)
}

func RemoveIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

func parseInput(filename string) (chunks []string) {

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		chunks = append(chunks, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return chunks
}
