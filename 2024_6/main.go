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

// function to count spaces visited
func counter(twodarray *[][]rune) int {
	counter := 1
	for _, line := range *twodarray {
		for _, r := range line {
			if r == 'O' || r == 'h' || r == 'v' || r == 'b' {
				counter++
			}
		}
	}
	return counter
}

func verticle(twodarray *[][]rune, cursor *Cursor) string {
	moves := 0
	switch cursor.Facing {
	case "yplus":
		for y := cursor.Y - 1; y >= 0; y-- {
			switch (*twodarray)[y][cursor.X] {
			case '#':
				if (*twodarray)[y+1][cursor.X] == 'h' {
					(*twodarray)[y+1][cursor.X] = 'b'
				} else {
					(*twodarray)[y+1][cursor.X] = 'v'
				}
				cursor.Y = y + 1
				return fmt.Sprintf("Cursor went up by %d moves.", moves)
			case 'v':
				return "loop"
			case 'b':
				return "loop"
			}
			moves++
			if (*twodarray)[y][cursor.X] == '.' {
				(*twodarray)[y][cursor.X] = 'O'
			}
		}
		cursor.Y = 0
		return fmt.Sprintf("Cursor went up by %d moves (to border).", moves)

	case "yminus":
		for y := cursor.Y + 1; y < len(*twodarray); y++ {
			switch (*twodarray)[y][cursor.X] {
			case '#':
				if (*twodarray)[y-1][cursor.X] == 'h' {
					(*twodarray)[y-1][cursor.X] = 'b'
				} else {
					(*twodarray)[y-1][cursor.X] = 'v'
				}
				cursor.Y = y - 1
				return fmt.Sprintf("Cursor went down by %d moves.", moves)
			case 'v':
				return "loop"
			case 'b':
				return "loop"
			}
			moves++
			if (*twodarray)[y][cursor.X] == '.' {
				(*twodarray)[y][cursor.X] = 'O'
			}
		}
		cursor.Y = len(*twodarray) - 1
		return fmt.Sprintf("Cursor went down by %d moves (to border).", moves)
	}
	return ""
}

func horizontal(twodarray *[][]rune, cursor *Cursor) string {
	moves := 0
	line := (*twodarray)[cursor.Y]
	switch cursor.Facing {
	case "xplus":
		for x := cursor.X + 1; x < len(line); x++ {
			switch line[x] {
			case '#':
				if (*twodarray)[cursor.Y][x-1] == 'v' {
					(*twodarray)[cursor.Y][x-1] = 'b'
				} else {
					(*twodarray)[cursor.Y][x-1] = 'h'
				}
				cursor.X = x - 1
				return fmt.Sprintf("Cursor went right by %d moves.", moves)
			case 'b':
				return "loop"
			case 'h':
				return "loop"
			}
			moves++
			if (*twodarray)[cursor.Y][x] == '.' {
				(*twodarray)[cursor.Y][x] = 'O'
			}
		}
		cursor.X = len(line) - 1
		return fmt.Sprintf("Cursor went right by %d moves (to border).", moves)
	case "xminus":
		for x := cursor.X - 1; x >= 0; x-- {
			switch line[x] {
			case '#':
				if (*twodarray)[cursor.Y][x+1] == 'v' {
					(*twodarray)[cursor.Y][x+1] = 'b'
				} else {
					(*twodarray)[cursor.Y][x+1] = 'h'
				}
				cursor.X = x + 1
				return fmt.Sprintf("Cursor went left by %d moves.", moves)
			case 'h':
				return "loop"
			case 'b':
				return "loop"
			}
			if line[x] == '#' {
				cursor.X = x + 1
				return fmt.Sprintf("Cursor went left by %d moves.", moves)
			}
			moves++
			if (*twodarray)[cursor.Y][x] == '.' {
				(*twodarray)[cursor.Y][x] = 'O'
			}
		}
		cursor.X = 0
		return fmt.Sprintf("Cursor went left by %d moves (to border).", moves)

	}
	return ""
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

func loop_check(x int, y int, data []byte) int {
	lines := to_lines(string(data))
	lines = lines[0 : len(lines)-1]
	lines[y][x] = '#'
	cursor := cursor_pos(lines)

	for (cursor.X != 0 && cursor.X != len(lines[0])-1) && (cursor.Y != 0 && cursor.Y != len(lines)-1) {
		switch cursor.Facing {
		case "yplus", "yminus":
			msg := verticle(&lines, &cursor)
			if msg == "loop" {
				return 1
			}
			cursor.TurnRight()
		case "xplus", "xminus":
			msg := horizontal(&lines, &cursor)
			if msg == "loop" {
				return 1
			}
			cursor.TurnRight()
		}
	}
	return 0
}

func main() {
	data, err := os.ReadFile("input")
	check(err)
	lines := to_lines(string(data))
	lines = lines[0 : len(lines)-1]

	// get initial cursor position
	cursor := cursor_pos(lines)
	moves := 0

	// part 1 solution:
	for (cursor.X != 0 && cursor.X != len(lines[0])-1) && (cursor.Y != 0 && cursor.Y != len(lines)-1) {
		moves++
		switch cursor.Facing {
		case "yplus", "yminus":
			msg := verticle(&lines, &cursor)
			fmt.Println(msg)
			fmt.Println("Coordinates (X, Y):", cursor.X, cursor.Y)
			cursor.TurnRight()
		case "xplus", "xminus":
			msg := horizontal(&lines, &cursor)
			fmt.Println(msg)
			fmt.Println("Coordinates (X, Y):", cursor.X, cursor.Y)
			cursor.TurnRight()
		}
	}
	ocount := counter(&lines)
	fmt.Println("Final state:", cursor.X, cursor.Y, cursor.Facing)
	fmt.Println("Spaces counted:", ocount)
	fmt.Println("Number of moves:", moves)

	// part 2 solution:
	lines2 := to_lines(string(data))
	loopCount := 0
	for y, line := range lines2 {
		for x, ch := range line {
			if ch == '.' {
				loopCount += loop_check(x, y, data)
			}
		}
	}
	fmt.Println("Number of options for creating a loop:", loopCount)
}
