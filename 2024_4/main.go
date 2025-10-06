package main

import (
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func reverse(s string) string {
	b := []byte(s)
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}

func to_lines(str string) []string {
	lines := strings.Split(str, "\n")
	return lines
}

func h_search(str string) int {
	count := 0
	for i := range str {
		if i+4 <= len(str) && (str[i:i+4] == "XMAS" || str[i:i+4] == "SAMX") {
			count++
		}
	}
	return count
}

func v_search(twodarray []string) int {
	count := 0
	for y, line := range twodarray {
		if y >= 3 {
			for x, letter := range line {
				if string(letter) == "X" || string(letter) == "S" {
					test_word := string(twodarray[y][x]) +
						string(twodarray[y-1][x]) +
						string(twodarray[y-2][x]) +
						string(twodarray[y-3][x])
					if test_word == "XMAS" || test_word == "SAMX" {
						count++
					}
				}
			}
		}
	}
	return count
}

func d_search(twodarray []string) int {
	count := 0
	for y, line := range twodarray {
		if y >= 3 {
			for x, letter := range line {
				if x >= 3 && (string(letter) == "X" || string(letter) == "S") {
					test_word := string(twodarray[y][x]) +
						string(twodarray[y-1][x-1]) +
						string(twodarray[y-2][x-2]) +
						string(twodarray[y-3][x-3])
					if test_word == "XMAS" || test_word == "SAMX" {
						count++
					}
				}
				if x+3 < len(line) && (string(letter) == "X" || string(letter) == "S") {
					test_word := string(twodarray[y][x]) +
						string(twodarray[y-1][x+1]) +
						string(twodarray[y-2][x+2]) +
						string(twodarray[y-3][x+3])
					if test_word == "XMAS" || test_word == "SAMX" {
						count++
					}
				}
			}
		}
	}
	return count
}

func find_xs(twodarray []string) int {
	count := 0
	for y, line := range twodarray {
		if y > 0 && y < len(twodarray)-1 {
			for x, letter := range line {
				if x > 0 && x < len(line)-1 {
					letter_str := string(letter)
					if letter_str == "A" {
						cross_1 := string(twodarray[y-1][x-1]) + letter_str + string(twodarray[y+1][x+1])
						cross_2 := string(twodarray[y-1][x+1]) + letter_str + string(twodarray[y+1][x-1])
						if (cross_1 == "MAS" || cross_1 == "SAM") && (cross_2 == "MAS" || cross_2 == "SAM") {
							count++
						}
					}
				}
			}
		}
	}
	return count
}

func main() {
	final_h_count := 0
	final_v_count := 0
	final_d_count := 0
	final_count := 0
	data, err := os.ReadFile("input")
	check(err)
	lines := to_lines(string(data))
	lines = lines[0 : len(lines)-1]
	for _, line := range lines {
		// counts horizontal matches
		h_count := h_search(line)
		final_h_count += h_count
		final_count += h_count
	}
	// counts vertical matches
	v_count := v_search(lines)
	final_v_count += v_count
	final_count += v_count

	// counts diagonal matches
	d_count := d_search(lines)
	final_d_count += d_count
	final_count += d_count

	// counts crossed MASes
	mass_count := find_xs(lines)

	fmt.Println("Horizontal:", final_h_count)
	fmt.Println("Vertical:", final_v_count)
	fmt.Println("Diagonal:", final_d_count)
	fmt.Println("Total:", final_count)
	fmt.Println("MAS count:", mass_count)

}
