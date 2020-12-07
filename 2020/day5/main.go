package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	f, err := os.Open("input")
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	prevSeat := -1
	a := make([]int, 0)

	// Input is a binary number
	// B and R is 1
	// F and L is 0
	// FFBFBFFRLL = 0010100100
	for scanner.Scan() {
		s := strings.ReplaceAll(scanner.Text(), "B", "1")
		s = strings.ReplaceAll(s, "F", "0")
		s = strings.ReplaceAll(s, "R", "1")
		s = strings.ReplaceAll(s, "L", "0")
		i, err := strconv.ParseInt(s, 2, 64)
		if err != nil {
			break
		}
		a = append(a, int(i))
	}
	sort.Ints(a)

	fmt.Printf("Part 1: %d\n", a[len(a)-1])

	for _, v := range a {
		if prevSeat != -1 && v != prevSeat+1 {
			fmt.Printf("Part 2: %d\n", v-1)
		}
		prevSeat = v
	}

}
