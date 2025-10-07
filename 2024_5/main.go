package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func convert(str string, divider string) [][]int {

	// split the string into a slice eg ["43|56" "23|78"]
	strs := strings.Split(str, "\n")

	// split each rule on the divider and add the result to a slice of slices [[43 56] [23 78]]
	var final_ints [][]int
	for _, rule := range strs {
		// makes strs ["43", "56"]
		strs := strings.Split(rule, divider)

		// makes ints 43, 56 and add to slice [43 56]
		var regrouped []int
		for _, str := range strs {
			n, _ := strconv.Atoi(str)
			regrouped = append(regrouped, n)
		}

		// adds them to final results slice [[43 56], ... ]
		final_ints = append(final_ints, regrouped)
	}
	return final_ints
}

func reorder(failure []int, rules [][]int) ([]int, int) {
	change_counter := 0
	for _, rule := range rules {
		if !(slices.Contains(failure, rule[0]) && slices.Contains(failure, rule[1])) {
			continue
		} else {
			var i0 int
			var i1 int
			for index, num := range failure {
				switch num {
				case rule[0]:
					i0 = index
				case rule[1]:
					i1 = index
				}

			}
			if i0 > i1 {
				val := failure[i0]
				failure = append(failure[:i0], failure[i0+1:]...)
				failure = slices.Insert(failure, i1, val)
				change_counter++
			}
		}
	}
	return failure, change_counter
}

func get_mid(s []int) int {
	mid := len(s) / 2
	mid_int := s[mid]
	return mid_int
}

func main() {
	data, err := os.ReadFile("input")
	check(err)
	divided := strings.Split(string(data), "\n\n")
	rules := convert(divided[0], "|")
	updates := convert(divided[1], ",")

	count_success := 0
	total := 0
	var failures [][]int
	for _, update := range updates {
		mid_int := get_mid(update)
		for _, rule := range rules {
			if !(slices.Contains(update, rule[0]) && slices.Contains(update, rule[1])) {
				continue
			} else {
				var i0 int
				var i1 int
				for index, num := range update {
					switch num {
					case rule[0]:
						i0 = index
					case rule[1]:
						i1 = index
					}

				}
				if i0 > i1 {
					mid_int = 0
				}
			}
		}
		total += mid_int
		if mid_int != 0 {
			count_success++
		} else if mid_int == 0 {
			failures = append(failures, update)
		}
	}

	sum_fixed := 0
	for _, failure := range failures {
		for {
			fixed, count := reorder(failure, rules)
			if count == 0 {
				mid := get_mid(fixed)
				sum_fixed += mid
				break
			}
		}
	}
	fmt.Println("Successes:", count_success)
	fmt.Println("Failures:", len(failures))
	fmt.Println("Total updates:", len(updates))
	fmt.Println("Answer to part 1:", total)
	fmt.Println("Answer to part 2:", sum_fixed)
}
