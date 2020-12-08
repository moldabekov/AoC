package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func unique(s string) int {
	alphabet := make(map[rune]bool)
	result := 0
	for _, i := range s {
		_, found := alphabet[i]
		if !found {
			result++
		}
		alphabet[i] = true
	}

	return result
}


// Set intersection
func overlap(s string, groupMembers int) int {
	alphabet := make(map[rune]int)
	result := 0

	for _, i :=range s {
		alphabet[i]++
	}

	for i:=0; i < len(s) / groupMembers; i++ {
		if alphabet[rune(s[i])] == groupMembers {
			result++
		}
	}

	return result
}

func main() {
	f, err := os.Open("input")
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	s := ""
	result := 0
	result2 := 0
	groupMembers := 0

	for scanner.Scan() {
		if scanner.Text() != "" {
			groupMembers++
			s += scanner.Text()
		} else {
			result += unique(s)
			result2 += overlap(s, groupMembers)

			groupMembers = 0
			s = ""
		}
	}

	fmt.Printf("Part 1: %d\n", result)
	fmt.Printf("Part 2: %d\n", result2)
}
