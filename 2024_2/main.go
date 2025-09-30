package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func to_lines(str string) [][]int {
	lines := strings.Split(str, "\n")
	var lines_ints [][]int
	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}
		line_ints := to_ints(line)
		lines_ints = append(lines_ints, line_ints)
	}
	return lines_ints
}
func to_ints(str string) []int {
	parts := strings.Fields(str)
	ints := make([]int, len(parts))
	for i, v := range parts {
		n, err := strconv.Atoi(v)
		check(err)
		ints[i] = n
	}
	return ints
}

func valid(line []int) bool {
	increasing := line[1] > line[0]
	for i := 0; i < len(line)-1; i++ {
		d := line[i+1] - line[i]
		if d == 0 || abs(d) > 3 {
			return false
		}
		if increasing && d < 0 || !increasing && d > 0 {
			return false
		}
	}
	return true
}

func pop_item(line []int, i int) []int {
	out := make([]int, 0, len(line)-1)
	out = append(out, line[:i]...)
	out = append(out, line[i+1:]...)
	return out
}

func dampen(line []int) bool {
	for i := range line {
		line_popped := pop_item(line, i)
		if valid(line_popped) {
			fmt.Println(line)
			fmt.Println(line_popped)
			return true
		}
	}
	return false
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	safe_reports := 0
	dampened_safe := 0
	var failures [][]int
	data, err := os.ReadFile("input")
	check(err)
	lines := string(data)
	lines_ints := to_lines(lines)
	fmt.Println("Total length:", len(lines_ints))
	for _, line := range lines_ints {
		if valid(line) {
			safe_reports++
		} else {
			failures = append(failures, line)
		}
	}
	for _, line := range failures {
		if dampen(line) {
			dampened_safe++
		}
	}

	shortest := 0
	for _, i := range lines_ints {
		if len(i) > shortest {
			shortest = len(i)
		}
	}
	fmt.Println("The shortest line is", shortest, "ints long.")
	fmt.Println("There were", safe_reports, "safe reports.")
	fmt.Println("There were", len(failures), "failures.")
	fmt.Println("The dampener adds", dampened_safe, "safe reports.")
	fmt.Println("Total safe reports:", safe_reports+dampened_safe)
}
