package main

import (
	"fmt"
)

const AlphabetSize  int = 26

func getAlphabetToNumberMap() [][]rune {
	phoneChar := [][]rune{
			{}, 
			{},
			{'a', 'b', 'c'},
			{'d', 'e', 'f'}, 
			{'g', 'h', 'i'},
			{'j', 'k', 'l'},
			{'m', 'n', 'o'},
			{'p', 'q', 'r', 's'},
			{'t', 'u', 'v'},
			{'w', 'x', 'y', 'z'},
	}
	return phoneChar
}

// Find possible strings match within the given number list 
// digits match to phone number digits
// example
func main() {

	// initialize the map with all the character

	
	//givenString := "98789872232"
	givenMatchStringList := []string{"abc", "dea", "ikl", "tmp", "flower", "atc"}

	//newList := make([]string, len(givenMatchStringList)) 
	c := []string{}
	//  convert given list of strings to number
	for index := 0 ; index < len(givenMatchStringList) ; index++ {
		a := ""
		for _, char :=  range givenMatchStringList[index] {
			 a =  a + string(char-'a')
			 fmt.Printf("string is %c %c\n", char, rune(char-'a'+'0'))
		}
		c = append(c, a)
	}

	fmt.Println(c)
}