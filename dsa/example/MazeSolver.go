package example

// coordinates as x & y index
type Point struct {
	X int
	Y int
}

type Maze struct {
	MazeArr []string // maze array
	Wall    string   // which char its the wall
	Start   Point    // entrance to the maze
	End     Point    // exit to the maze
}

var dir = [4][2]int{
	{1, 0},
	{0, -1},
	{0, 1},
	{-1, 0},
}

func (m *Maze) walk(curr Point, seen [][]bool, path *[]Point) bool {
	// base cases for exiting

	// 1. off the MazeArr
	if curr.X < 0 || curr.X >= len(m.MazeArr[0]) || curr.Y < 0 || curr.Y >= len(m.MazeArr) {
		return false
	}

	// 2. on a wall
	if string(m.MazeArr[curr.Y][curr.X]) == m.Wall {
		return false
	}

	// 3. found the End
	if curr.X == m.End.X && curr.Y == m.End.Y {
		*path = append(*path, curr)
		return true
	}

	// 4. already seen
	if seen[curr.Y][curr.X] {
		return false
	}

	// check for path in all the directions
	// pre
	seen[curr.Y][curr.X] = true
	*path = append(*path, curr)

	// recurse
	for i := 0; i < len(dir); i++ {
		X, Y := dir[i][0], dir[i][1]

		newCurr := Point{
			X: curr.X + X,
			Y: curr.Y + Y,
		}

		if m.walk(newCurr, seen, path) {
			return true
		}
	}

	// post
	*path = (*path)[:len(*path)-1]

	return false
}

func (m *Maze) Solve() []Point {
	seen := [][]bool{}
	path := []Point{}

	for i := 0; i < len(m.MazeArr); i++ {
		seen = append(seen, make([]bool, len(m.MazeArr[0])))
	}

	m.walk(m.Start, seen, &path)

	return path
}
