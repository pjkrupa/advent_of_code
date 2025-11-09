package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

type Point struct {
	X, Y int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// check to see if point is already recorded as an antinode
func check_exists(p Point, seen map[Point]struct{}) bool {
	_, exists := seen[p]
	return exists
}

func to_lines(str string) [][]rune {
	lines := strings.Split(str, "\n")
	rs := make([][]rune, len(lines))
	for i, line := range lines {
		rs[i] = []rune(line)
	}
	return rs
}

func edge_check(p Point, lines [][]rune) bool {
	switch {
	case p.X > len(lines[0])-1:
		return false
	case p.X < 0:
		return false
	case p.Y > len(lines)-1:
		return false
	case p.Y < 0:
		return false
	default:
		return true
	}
}

func calc_antinodes(p2, p3 Point) (Point, Point) {
	x1, y1 := float64(p2.X), float64(p2.Y)
	x2, y2 := float64(p3.X), float64(p3.Y)

	dist := math.Hypot(x2-x1, y2-y1)
	if dist == 0 {
		return Point{}, Point{}
	}

	dx := (x2 - x1) / dist
	dy := (y2 - y1) / dist

	p1 := Point{X: int(math.Round(x1 - dx*dist)), Y: int(math.Round(y1 - dy*dist))}
	p4 := Point{X: int(math.Round(x2 + dx*dist)), Y: int(math.Round(y2 + dy*dist))}

	return p1, p4
}

func antinode_check(lines [][]rune, p Point, seen map[Point]struct{}) {
	if edge_check(p, lines) && !check_exists(p, seen) {
		seen[p] = struct{}{}
	}
}

func search(lines [][]rune, p Point, r rune, seen map[Point]struct{}) {
	// search up from point
	for y := p.Y - 1; y >= 0; y-- {
		for x, n := range lines[y] {
			if n == r {
				p2 := Point{X: x, Y: y}
				a1, a2 := calc_antinodes(p, p2)
				antinode_check(lines, a1, seen)
				antinode_check(lines, a2, seen)
			}
		}
	}

	// search down from point
	for y := p.Y + 1; y < len(lines); y++ {
		for x, n := range lines[y] {
			if n == r {
				p2 := Point{X: x, Y: y}
				a1, a2 := calc_antinodes(p, p2)
				antinode_check(lines, a1, seen)
				antinode_check(lines, a2, seen)
			}
		}
	}
}

func main() {
	data, err := os.ReadFile("input")
	check(err)
	lines := to_lines(string(data))
	lines = lines[0 : len(lines)-1]
	seen := make(map[Point]struct{})
	// main loop: loop through each line
	for y, line := range lines {
		for x, r := range line {
			if r != '.' {
				p := Point{X: x, Y: y}
				// function that searches for antennae, returns count of antinodes
				search(lines, p, r, seen)
			}
		}
	}
	count := len(seen)
	fmt.Println("Antinodes:", seen)
	fmt.Println("Total antinodes:", count)
}
