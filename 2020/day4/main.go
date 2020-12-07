package main

import (
	"bufio"
	"encoding/hex"
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

func validate1(passport map[string]string) bool {
	if _, ok := passport["byr"]; !ok {
		return false
	} else if _, ok := passport["iyr"]; !ok {
		return false
	} else if _, ok := passport["eyr"]; !ok {
		return false
	} else if _, ok := passport["hgt"]; !ok {
		return false
	} else if _, ok := passport["hcl"]; !ok {
		return false
	} else if _, ok := passport["ecl"]; !ok {
		return false
	} else if _, ok := passport["pid"]; !ok {
		return false
	}
	return true
}

func validate2(passport map[string]string) bool {
	byr, _ := strconv.Atoi(passport["byr"])
	iyr, _ := strconv.Atoi(passport["iyr"])
	eyr, _ := strconv.Atoi(passport["eyr"])

	// birth year
	if 1920 > byr || byr > 2002 {
		return false
	}

	// issue year
	if 2010 > iyr || iyr > 2020 {
		return false
	}

	// expiration year
	if 2020 > eyr || eyr > 2030 {
		return false
	}

	// height
	// Try to split string by "cm"
	hgt := strings.Split(passport["hgt"], "cm")
	if len(hgt) > 1 {
		h, _ := strconv.Atoi(hgt[0])
		if 150 > h || h > 193 {
			return false
		}
	// if it fails, then split by "in"
	} else {
		hgt := strings.Split(passport["hgt"], "in")
		h, _ := strconv.Atoi(hgt[0])
		if 59 > h || h > 76 {
			return false
		}
	}

	// hair color
	// try to decode hex without #
	if _, err := hex.DecodeString(passport["hcl"][1:]); err != nil {
		return false
	}

	// eye color
	switch passport["ecl"] {
	case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
		// continue
	default:
		return false
	}

	// passport id
	if (len(passport["pid"]) == 9) {
		if _, err := strconv.Atoi(passport["pid"]); err != nil {
			return false
		}
	} else {
		return false
	}
	return true
}

// I intentionally didn't use regexp, since regexps are slow as hell in Go.
// As well we didn't use Struct, because it'd take time to implement parser.
// Edge case: EOF of input should be double blank line. Lazy to handle it.
func main() {
	f, err := os.Open("input")
	check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)

	passport := make(map[string]string)
	// Passports fields are:
	// byr (Birth Year) [int]
	// iyr (Issue Year) [int]
	// eyr (Expiration Year) [int]
	// hgt (Height) [string]
	// hcl (Hair Color) [hex]
	// ecl (Eye Color)  [string]
	// pid (Passport ID) [int]
	// cid (Country ID)

	var result int64
	var result2 int64

	for scanner.Scan() {
		if scanner.Text() != "" {
			s := strings.Split(scanner.Text(), " ") // extracting fields 1by1
			for i := 0; i < len(s); i++ {
				passport[strings.Split(s[i], ":")[0]] = strings.Split(s[i], ":")[1] // mapping KVs to hashmap
			}
		} else {
			if len(passport) > 6 && validate1(passport) {
				result++

				// lets validate passports that look good at a first glance
				if validate2(passport) {
					result2++
				}
			}
			passport = make(map[string]string)
		}
	}

	fmt.Printf("Part 1: %d\n", result)
	fmt.Printf("Part 2: %d\n", result2)
}
