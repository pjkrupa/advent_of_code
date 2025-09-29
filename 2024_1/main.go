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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func get_min(lines []int) int {
	m := 0
	for _, val := range lines {
		if val < m {
			m = val
		}
	}
	return m
}

func cast_int(str string) int {
	i, _ := strconv.Atoi(str)
	return i
}

func get_ints(line string) (int, int) {
	split_line := strings.Split(line, "   ")

	item0 := strings.TrimSpace(split_line[0])
	item1 := strings.TrimSpace(split_line[1])

	item0_int := cast_int(item0)
	item1_int := cast_int(item1)
	return item0_int, item1_int
}

func part2_calc(list0, list1 []int) int {
	final_sum := 0
	for _, val0 := range list0 {
		count := 0
		for _, val1 := range list1 {
			if val0 == val1 {
				count += 1
			}
		}
		product := val0 * count
		final_sum += product
	}
	return final_sum
}

func main() {
	var list0 []int
	var list1 []int

	dat, err := os.ReadFile("input")
	check(err)
	lines := strings.Split(string(dat), "\n")
	for _, line := range lines {
		var int0 int
		var int1 int
		if line == "" {
			int0 = 0
			int1 = 0
		} else {
			int0, int1 = get_ints(line)
		}

		list0 = append(list0, int0)
		list1 = append(list1, int1)

	}
	slices.Sort(list0)
	slices.Sort(list1)
	// count0 := len(list0)
	// count1 := len(list1)

	var total int

	for i := 0; i <= 1000; i++ {
		diff := list0[i] - list1[i]
		result := abs(diff)
		total += result
	}

	total2 := part2_calc(list0, list1)

	fmt.Println(total)
	fmt.Println(total2)
	// fmt.Println("There are", count0, "items in list0.")
	// fmt.Println("There are", count1, "items in list1.")
	// fmt.Println("Last item in list0:", list0[len(list0)-1])
	// fmt.Println("Last item in list1:", list1[len(list1)-1])
}
