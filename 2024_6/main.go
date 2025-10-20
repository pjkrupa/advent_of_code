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

// - X find the location of the pointer
// - X create struct for pointer direction (yplus, yminus, xplus, xminus)
// - set position values as "yplus" until #
// - turn based on pointer direction (case switch), update enum value
// - repeat until outide range
// - count Xs

type Cursor struct {
	X      int
	Y      int
	Facing string
}

func (d *Cursor) TurnRight() {
	switch d.Facing {
	case "xplus":
		d.Facing = "yminus"
	case "yminus":
		d.Facing = "xminus"
	case "xminus":
		d.Facing = "yplus"
	case "yplus":
		d.Facing = "xplus"
	}
}

// // change this to work with runes instead of strings
func yplus(twodarray *[][]rune, cursor *Cursor) string {

	for y := len(*twodarray) - cursor.Y; y >= 0; y-- {
		if (*twodarray)[y][cursor.X] == '#' {
			cursor.Y = y
		} else {
			(*twodarray)
		}
	}
}

// finds the initial position of the cursor, instantiates it, and returns it
func cursor_pos(twodarray [][]rune) Cursor {
	cur := Cursor{X: 0, Y: 0, Facing: "yplus"}
	for y, line := range twodarray {
		for x, ch := range line {
			if ch == '^' {
				cur.X = x
				cur.Y = y
				return cur
			}
		}
	}
	return cur
}

func to_lines(str string) [][]rune {
	lines := strings.Split(str, "\n")
	rs := make([][]rune, len(lines))
	for i, line := range lines {
		rs[i] = []rune(line)
	}
	return rs
}

func main() {
	data, err := os.ReadFile("input")
	check(err)
	lines := to_lines(string(data))
	lines = lines[0 : len(lines)-1]
	cursor := cursor_pos(lines)

	fmt.Println(cursor.X, cursor.Y, cursor.Facing)
	fmt.Println(lines[cursor.Y])
}
