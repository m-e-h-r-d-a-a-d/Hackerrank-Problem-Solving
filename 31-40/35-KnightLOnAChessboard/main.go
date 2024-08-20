package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x, y int32
}

func bfs(seen [][]int32, lenght int32, step int32, p point, s point) int32 {

	if p.x == lenght && p.y == lenght {
		return step
	}

	if p.x < 0 || p.y < 0 || p.x > lenght || p.y > lenght {
		return -1
	}

	step++
	if seen[p.x][p.y] != 0 && seen[p.x][p.y] <= step {
		return -1
	}

	seen[p.x][p.y] = step

	min := checkMin(bfs(seen, lenght, step, point{x: p.x - s.x, y: p.y - s.y}, s), -1)
	min = checkMin(bfs(seen, lenght, step, point{x: p.x + s.x, y: p.y - s.y}, s), min)
	min = checkMin(bfs(seen, lenght, step, point{x: p.x - s.x, y: p.y + s.y}, s), min)
	min = checkMin(bfs(seen, lenght, step, point{x: p.x + s.x, y: p.y + s.y}, s), min)
	min = checkMin(bfs(seen, lenght, step, point{x: p.x - s.y, y: p.y - s.x}, s), min)
	min = checkMin(bfs(seen, lenght, step, point{x: p.x + s.y, y: p.y - s.x}, s), min)
	min = checkMin(bfs(seen, lenght, step, point{x: p.x - s.y, y: p.y + s.x}, s), min)
	min = checkMin(bfs(seen, lenght, step, point{x: p.x + s.y, y: p.y + s.x}, s), min)

	return min
}

func checkMin(v int32, min int32) int32 {
	if v == -1 {
		return min
	}

	if min == -1 {
		return v
	}

	if v < min {
		return v
	}

	return min
}

func knightlOnAChessboard(n int32) [][]int32 {

	out := make([][]int32, n-1)
	for i := int32(1); i < n; i++ {
		out[i-1] = make([]int32, n-1)
		for j := int32(1); j < n; j++ {
			seen := make([][]int32, n)
			for k := int32(0); k < n; k++ {
				seen[k] = make([]int32, n)
			}

			out[i-1][j-1] = bfs(seen, n-1, 0, point{x: 0, y: 0}, point{x: i, y: j})
		}
	}

	return out
}

func main() {
	osFile, _ := os.Open("TestCase0.txt")
	reader := bufio.NewReader(osFile)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	result := knightlOnAChessboard(n)

	fmt.Println(result)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
