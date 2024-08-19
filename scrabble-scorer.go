package main

import (
	"fmt"
	"strings"
)

var oldPointStructure = map[int][]string{
	1:  {"A", "E", "I", "O", "U", "L", "N", "R", "S", "T"},
	2:  {"D", "G"},
	3:  {"B", "C", "M", "P"},
	4:  {"F", "H", "V", "W", "Y"},
	5:  {"K"},
	8:  {"J", "X"},
	10: {"Q", "Z"},
}

func OldScrabbleScore(word string) string {
	word = strings.ToUpper(word)
	letterPoints := ""

	for _, ch := range word {
		// fmt.Printf("%c\n", ch)
		lettercheck := string(ch)
		// fmt.Println(lettercheck)
		for key, value := range oldPointStructure {
			for _, l := range value {
				if l == lettercheck {
					letterPoints += fmt.Sprintf("Points for '%s': %d\n", lettercheck, key)
				}
			}
		}
	}

	return letterPoints
}
func initialPrompt() string {
	var userWord string
	fmt.Println("Let's play some scrabble! Enter a word here: ")
	fmt.Scan(&userWord)

	return userWord
}

var newPointStructure = transform(oldPointStructure)

var simpleScore = func(word string) int {
	pointTotal := len(word)
	return pointTotal

}
var vowelBonusScore = func(word string) int {

	word = strings.ToUpper(word)
	var pointTotal int = 0
	vowels := "AEIOU"
	for _, letter := range word {
		letterCheck := string(letter)
		if strings.Contains(vowels, letterCheck) {
			pointTotal += 3
		} else {
			pointTotal += 1
		}
	}
	return pointTotal
}
var scrabbleScore = func(word string) int {

	var pointTotal int = 0
	for _, ch := range word {
		letter := strings.ToUpper(string(ch))
		pointTotal += newPointStructure[letter]
	}

	return pointTotal
}

type scoreObject struct {
	name        string
	description string
	algorithm   func(string) int
}

var algorithmSlice = []scoreObject{
	{"Simple Score", "1 point per letter", simpleScore},
	{"Vowel Bonus", "3 points for vowels, 1 point for consonants", vowelBonusScore},
	{"Scrabble Score", "Traditional scrabble point values", scrabbleScore},
}

func scoreSelect() scoreObject {
	var userChoice int
	fmt.Println("Great word! Choose a scoring algorithm to use")

	for index, alg := range algorithmSlice {
		fmt.Printf("\n%v: \nName:%v\nDescription:%v\n", index, alg.name, alg.description)
	}
	fmt.Scan(&userChoice)

	return algorithmSlice[userChoice]
}

func transform(ogStructure map[int][]string) map[string]int {
	var newPointStructure = map[string]int{}
	//outer loop iterates over old point structure
	for key, value := range oldPointStructure {
		// loop over value array
		for _, letter := range value {
			if _, ok := newPointStructure[letter]; !ok {
				newPointStructure[letter] = key
			}
		}
	}

	return newPointStructure
}

func RunProgram() {

	var userWord string = initialPrompt()
	// fmt.Println(vowelBonusScore(userWord))
	var userAlgo scoreObject = scoreSelect()
	var points int = userAlgo.algorithm(userWord)
	fmt.Printf("your word %v earned %v points using the %v scoring algorithm", userWord, points, userAlgo.name)
}
