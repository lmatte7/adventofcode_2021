package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {

	template, pairs := parseInput("input.txt")
	steps := 10

	// fmt.Println(template)

	// for key, value := range pairs {
	// 	fmt.Printf("Key: %s, Value %s\n", key, value)
	// }
	letterFrequencies := make(map[string]int)

	newTemplate := template
	for i := 0; i < steps; i++ {
		if i > 0 {
			template = newTemplate
		}

		newTemplateIndex := 1
		for key, letter := range strings.Split(template, "") {
			if key < len(template)-1 {
				currentPair := letter + string(template[key+1])

				newTemplate = newTemplate[:newTemplateIndex] + pairs[currentPair] + newTemplate[newTemplateIndex:]
				newTemplateIndex += 2
			}
		}
	}

	for _, letter := range strings.Split(newTemplate, "") {
		letterFrequencies[letter]++
	}

	highest := 0
	lowest := 0
	i := 0
	for _, frequency := range letterFrequencies {
		if i == 0 {
			lowest = frequency
		}

		if highest < frequency {
			highest = frequency
		}

		if lowest > frequency {
			lowest = frequency
		}

		i++
	}

	fmt.Println(len(newTemplate))
	fmt.Println(highest - lowest)
}

func parseInput(filename string) (template string, pairs map[string]string) {

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	pairs = make(map[string]string)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		line := scanner.Text()
		isPolymerPair, err := regexp.MatchString("->", line)
		if err != nil {
			panic(err)
		}
		if isPolymerPair {
			pair := strings.Split(line, " -> ")
			pairs[pair[0]] = pair[1]
		} else {
			if len(line) > 0 {
				template = line
			}
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return
}
