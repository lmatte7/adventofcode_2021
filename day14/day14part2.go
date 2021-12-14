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
	steps := 40

	// fmt.Println(template)

	// for key, value := range pairs {
	// 	fmt.Printf("Key: %s, Value %s\n", key, value)
	// }
	pairFrequencies := make(map[string]int)
	letterFrequencies := make(map[string]int)

	for i := 0; i < len(template); i++ {

		if i < len(template)-1 {
			pairFrequencies[string(template[i])+string(template[i+1])]++
		}
	}

	first := template[0]
	letterFrequencies[string(first)]++

	for i := 0; i < steps; i++ {
		newFrequencies := copyFrequencies(pairFrequencies)

		for rulePair, addedLetter := range pairs {

			ruleFrequency := pairFrequencies[rulePair]
			if ruleFrequency > 0 {
				newFrequencies[rulePair] -= ruleFrequency

				newFrequencies[string(rulePair[0])+addedLetter] += ruleFrequency
				newFrequencies[addedLetter+string(rulePair[1])] += ruleFrequency
			}
		}

		pairFrequencies = newFrequencies
	}

	for pair, frequency := range pairFrequencies {
		letterFrequencies[string(pair[1])] += frequency
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

	fmt.Println(pairFrequencies)
	fmt.Println(letterFrequencies)
	fmt.Println(highest - lowest)
}

func copyFrequencies(pairFrequencies map[string]int) map[string]int {
	newMap := make(map[string]int)
	for k, v := range pairFrequencies {
		newMap[k] = v
	}
	return newMap
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
