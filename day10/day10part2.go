package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {

	chunks := parseInput("input.txt")

	badCharacterScore := 0
	finalScores := make([]int, 0)
	for _, chunk := range chunks {
		finishingCharacterScore := 0

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
					// fmt.Printf("Bad Character Found: %s\n", characters[i])
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
					// fmt.Printf("Bad Character Found: %s\n", characters[i])
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
					// fmt.Printf("Bad Character Found: %s\n", characters[i])
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
					// fmt.Printf("Bad Character Found: %s\n", characters[i])
					break
				}
			}

		}
		if valid {
			// fmt.Printf("Remaining Characters: %v\n", closingCharacters)
			for i := (len(closingCharacters) - 1); i >= 0; i-- {
				// fmt.Printf("%s", closingCharacters[i])
				finishingCharacterScore *= 5
				switch closingCharacters[i] {
				case "(":
					finishingCharacterScore += 1
				case "[":
					finishingCharacterScore += 2
				case "{":
					finishingCharacterScore += 3
				case "<":
					finishingCharacterScore += 4
				}
			}
			// fmt.Println()
			finalScores = append(finalScores, finishingCharacterScore)
		}
		// fmt.Printf("String %s is %v\n", chunk, valid)
	}
	sort.Ints(finalScores[:])
	fmt.Printf("Character Score: %v\n", finalScores)

	middle := (len(finalScores) / 2)

	fmt.Printf("Middle result: %d", finalScores[middle])
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
