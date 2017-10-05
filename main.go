package main

import "fmt"
import (
	"strings"
	"os"
	"bufio"
	"os/exec"
)

type Position struct {
	X int
	Y int
}

func (me *Position) move(matrix [][]string, dir string, ) (bool, [][]string) {

	nextPosition := Position{me.X, me.Y}
	moved := false

	switch dir {
	case "w":
		nextPosition.Y--
	case "d":
		nextPosition.X++
	case "a":
		nextPosition.X--
	case "s":
		nextPosition.Y++
	}

	if nextPosition.Y >= 0 && nextPosition.X >= 0 &&
		nextPosition.Y < len(matrix) && nextPosition.X < len(matrix[nextPosition.Y]) &&
		matrix[nextPosition.Y][nextPosition.X] != "#" {
		matrix[nextPosition.Y][nextPosition.X] = "*"
		matrix[me.Y][me.X] = " "
		me.X = nextPosition.X
		me.Y = nextPosition.Y
		moved = true
	}

	return moved, matrix
}

func main() {

	m1 := []string{
		"######",
		"#  * #",
		"# #  #",
		"#A## #",
		"# ^# #",
		"#### #",
		"#a   #",
		"######",
	}

	m2 := []string{
		"#############",
		"       ^#*   ",
		"#############",
	}

	m3 := []string{
		"########",
		"#  #  ^#",
		"#*  #  #",
		"########",
	}

	m4 := []string{
		"^XcYbZa * AzByCx#",
	}

	_, _, _, _ = m1, m2, m3, m4

	dungeonBlaster(m1)
	// dungeonBlaster(m2)
	// dungeonBlaster(m3)
	// dungeonBlaster(m4)

}

func dungeonBlaster(m []string) bool {

	const (
		xVisionRange = 2
		yVisionRange = 2
	)

	matrix, me := makeMatrix(m)

	scanner := bufio.NewScanner(os.Stdin)
Render:
	for scanner.Scan() {
		input := strings.Split(scanner.Text(), "")
		for _, command := range input {
			switch command {
			case "w":
				fallthrough
			case "d":
				fallthrough
			case "a":
				fallthrough
			case "s":
				_, matrix := me.move(matrix, command)
				render(matrix)
			case "q":
				break Render
			}
		}
	}

	return true
}

func makeMatrix(m []string) ([][]string, Position) {
	matrix := make([][]string, len(m))
	var me Position
	for y, row := range m {
		matrix[y] = strings.Split(row, "")
		for x, block := range matrix[y] {
			if block == "*" {
				me = Position{x, y}
			}
		}
	}
	return matrix, me
}

func render(matrix [][]string) string {
	dungeonView := "\n"
	for _, row := range matrix {
		dungeonView += strings.Join(row, "") + "\n"
	}

	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()

	fmt.Println(dungeonView)

	return dungeonView
}
