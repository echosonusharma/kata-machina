package dsa_test

import (
	"fmt"
	"strings"
	"testing"

	dsa "github.com/echosonusharma/kata-machina/dsa/example"
)

func TestMazeSolver(t *testing.T) {
	maze := []string{
		"## #",
		"#  #",
		"# ##",
	}

	mazeStart := dsa.Point{
		X: 2,
		Y: 0,
	}

	mazeEnd := dsa.Point{
		X: 1,
		Y: 2,
	}

	m := &dsa.Maze{
		MazeArr: maze,
		Wall:    "#",
		Start:   mazeStart,
		End:     mazeEnd,
	}

	r := m.Solve()

	mazePathExpected := []dsa.Point{
		{X: 2, Y: 0},
		{X: 2, Y: 1},
		{X: 1, Y: 1},
		{X: 1, Y: 2},
	}

	rGot, got := drawPath(maze, r)
	rExp, exp := drawPath(maze, mazePathExpected)
	printRes(rGot, rExp, got, exp)
	if got != exp {
		t.Error("did'nt get the expected path for maze")
	}

	//

	maze1 := []string{
		"xxxxxxxxxx x",
		"x        x x",
		"x        x x",
		"x xxxxxxxx x",
		"x          x",
		"x xxxxxxxxxx",
	}

	mazeStart1 := dsa.Point{
		X: 10,
		Y: 0,
	}

	mazeEnd1 := dsa.Point{
		X: 1,
		Y: 5,
	}

	m1 := &dsa.Maze{
		MazeArr: maze1,
		Wall:    "x",
		Start:   mazeStart1,
		End:     mazeEnd1,
	}

	r1 := m1.Solve()

	mazePathExpected1 := []dsa.Point{
		{X: 10, Y: 0},
		{X: 10, Y: 1},
		{X: 10, Y: 2},
		{X: 10, Y: 3},
		{X: 10, Y: 4},
		{X: 9, Y: 4},
		{X: 8, Y: 4},
		{X: 7, Y: 4},
		{X: 6, Y: 4},
		{X: 5, Y: 4},
		{X: 4, Y: 4},
		{X: 3, Y: 4},
		{X: 2, Y: 4},
		{X: 1, Y: 4},
		{X: 1, Y: 5},
	}

	rGot1, got1 := drawPath(maze1, r1)
	rExp1, exp1 := drawPath(maze1, mazePathExpected1)
	printRes(rGot1, rExp1, got1, exp1)
	if got1 != exp1 {
		t.Error("did'nt get the expected path for maze1")
	}

}

func drawPath(maze []string, path []dsa.Point) (rawM [][]string, result string) {
	m := [][]string{}

	for _, v := range maze {
		m = append(m, strings.Split(v, ""))
	}

	for i := 0; i < len(path); i++ {
		x, y := path[i].X, path[i].Y
		m[y][x] = string('*')
	}

	r := ""

	for _, v := range m {
		r += strings.Join(v, "")
	}

	return m, r
}

func printRes(rGot, rExp [][]string, got, exp string) {
	fmt.Println("--- got ---")
	for _, v := range rGot {
		fmt.Println(v)
	}
	fmt.Println("--- expected ---")
	for _, v := range rExp {
		fmt.Println(v)
	}
	fmt.Println("---------")
	fmt.Println("got - ", got)
	fmt.Println("exp - ", exp)
	fmt.Println("---------")
}
