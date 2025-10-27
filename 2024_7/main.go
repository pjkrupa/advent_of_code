package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Line struct {
	test_int int
	nums     []int
	seqs     []string
}

// takes a sequence (eg: "aamm") and checks whether, when applied to the numbers, the total matches the test_int
func calculate(str string, test_nums []int, test_int int) int {
	nums := append([]int(nil), test_nums...)
	for i, char := range str {
		// fmt.Println(nums)
		// fmt.Println(str)
		if i != len(nums)-1 {
			switch char {
			case '*':
				// fmt.Println(nums[i] * nums[i+1])
				nums[i+1] = nums[i] * nums[i+1]
			case '+':
				// fmt.Println(nums[i] + nums[i+1])
				nums[i+1] = nums[i] + nums[i+1]
			case '|':
				s := strconv.Itoa(nums[i]) + strconv.Itoa(nums[i+1])
				nums[i+1], _ = strconv.Atoi(s)
			}
		}
		if nums[len(nums)-1] == test_int {
			return test_int
		}
	}
	return 0
}

func make_sequences(nums []int) []string {
	outcomes := []string{""}
	n := len(nums) - 1
	for range n {
		var newOutcomes []string
		for _, s := range outcomes {
			newOutcomes = append(newOutcomes, s+"+")
			newOutcomes = append(newOutcomes, s+"*")
			newOutcomes = append(newOutcomes, s+"|")
		}
		outcomes = newOutcomes
	}
	return outcomes
}

func make_line(str string) Line {
	sl_ints := []int{}
	sl := strings.Split(str, ":")
	t_int, _ := strconv.Atoi(sl[0])
	sl_strs := strings.Split(strings.TrimSpace(sl[1]), " ")
	for _, s := range sl_strs {
		i, _ := strconv.Atoi(s)
		sl_ints = append(sl_ints, i)
	}
	sequences := make_sequences(sl_ints)
	line := Line{test_int: t_int, nums: sl_ints, seqs: sequences}
	return line
}

func get_lines(str string) []Line {
	var lines []Line
	sl_strs := strings.Split(str, "\n")
	sl_strs = sl_strs[0 : len(sl_strs)-1]
	for _, s := range sl_strs {
		line := make_line(s)
		lines = append(lines, line)
	}
	return lines
}
func main() {
	data, _ := os.ReadFile("input")
	total := 0
	lines := get_lines(string(data))
	for _, line := range lines {
		fmt.Println(line.test_int, ":", line.nums)
		for _, seq := range line.seqs {
			nums := line.nums
			result := calculate(seq, nums, line.test_int)
			total += result
			if result != 0 {
				fmt.Println("found one:", result)
				fmt.Println("for:", seq)
				break
			}
		}
	}
	fmt.Println("grand total:", total)
}
