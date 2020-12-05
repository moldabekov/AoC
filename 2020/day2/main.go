package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func validate1(min int, max int, char string, password string) bool {
	x := strings.Count(password, char)
	if min <= x && x <= max {
		return true
	}
	return false
}

func validate2(min int, max int, char string, password string) bool {
	if password[min-1:min] == char && password[max-1:max] != char ||
		password[min-1:min] != char && password[max-1:max] == char {
		return true
	}
	return false
}

func main() {
	f, err := os.Open("input")
	check(err)
	defer f.Close()

	valid1 := 0
	valid2 := 0

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		count := strings.Split(scanner.Text(), "-") // get slice of strs with min-max
		min, _ := strconv.Atoi(count[0])                // convert to int for further use
		max, _ := strconv.Atoi(count[1])
		scanner.Scan()
		char := scanner.Text()[0:1]                     // get char
		scanner.Scan()
		password := scanner.Text()                      // get password string

		if validate1(min, max, char, password) {
			valid1 = valid1 + 1
		}

		if validate2(min, max, char, password) {
			valid2 = valid2 + 1
		}
	}

	fmt.Printf("Part 1: %d\n", valid1)
	fmt.Printf("Part 2: %d\n", valid2)

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
