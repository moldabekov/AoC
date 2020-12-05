package main

import (
    "bufio"
    "fmt"
    "os"
    "sort"
    "strconv"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func findYearsPart1(a []int) (int) {
    for _, year := range a {
        year2 := 2020 - year
        if j := sort.SearchInts(a, 2020-year); year2 == a[j] {
            return year * year2
        }
    }
    return -1
}

func findYearsPart2(a []int) (int) {
    for i, _ := range a {
        l := i+1
        r := len(a) - 1
        for l < r {
            if a[i] + a[l] + a[r] == 2020 {
                return a[i] * a[l] * a[r]
            }
            if a[i] + a[l] + a[r] < 2020 {
                l++
            }
            if a[i] + a[l] + a[r] > 2020 {
                r--
            }
        }
    }
    return -1
}

func main() {
    f, err := os.Open("input")
    check(err)
    defer f.Close()

    a := make([]int, 0)

    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        if i, err := strconv.Atoi(scanner.Text()); err == nil {
            a = append(a, i)
        }
    }

    sort.Ints(a)
    fmt.Printf("Part 1: %d\n", findYearsPart1(a))
    fmt.Printf("Part 2: %d\n", findYearsPart2(a))
}
