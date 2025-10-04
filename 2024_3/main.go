package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func calculate(str string) int {
	re_nums := regexp.MustCompile(`\d+,\d+`)
	str_clean := re_nums.FindAllString(str, 1)
	strs := strings.Split(str_clean[0], ",")
	// fmt.Println("String slice:", strs)
	var nums []int
	for _, str := range strs {
		num, _ := strconv.Atoi(str)
		nums = append(nums, num)
	}
	// fmt.Println("Ints slice:", nums)
	final_int := nums[0] * nums[1]
	return final_int
}

func find_matches(str string) []string {
	re := regexp.MustCompile(`\bmul\(\d+,\d+\)`)
	matches := re.FindAllString(str, -1)
	return matches
}

func make_sections(str string) []string {
	sections := strings.Split(str, "do")
	return sections
}

func main() {
	data, err := os.ReadFile("input")
	check(err)
	my_string := string(data)
	matches := find_matches(my_string)
	total := 0
	for _, str := range matches {
		prod := calculate(str)
		total += prod
	}
	fmt.Println("Final result:", total)

	do_total := 0
	do_sections := make_sections(my_string)
	for _, section := range do_sections {
		// fmt.Println("---------------------------------------------------------------")
		// fmt.Println(section)
		if !strings.HasPrefix(section, "n't") {
			do_matches := find_matches(section)
			for _, str := range do_matches {
				do_prod := calculate(str)
				do_total += do_prod
			}
		}
	}
	fmt.Println("Do result:", do_total)
	back_together := strings.Join(do_sections, "do")
	fmt.Println("Initial length:", len(my_string))
	fmt.Println("Final length:", len(back_together))
}
