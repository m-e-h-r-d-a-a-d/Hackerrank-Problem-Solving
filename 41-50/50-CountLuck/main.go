package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const (
	max = int32(1000000000)
)

func findKey(m [][]rune, x int32, y int32, w int32, h int32, s int32) (int32, bool) {
	if x < 0 || x >= w || y < 0 || y >= h || m[x][y] == 'X' {
		return max, false
	}

	if m[x][y] == '*' {
		return s, true
	}

	counter := 0
	if x > 0 && (m[x-1][y] == '.' || m[x-1][y] == '*') {
		counter++
	}

	if y > 0 && (m[x][y-1] == '.' || m[x][y-1] == '*') {
		counter++
	}

	if x < w-1 && (m[x+1][y] == '.' || m[x+1][y] == '*') {
		counter++
	}

	if y < h-1 && (m[x][y+1] == '.' || m[x][y+1] == '*') {
		counter++
	}

	if counter > 1 {
		s++
	}

	m[x][y] = 'X'

	r := max
	f := false

	if rn, fn := findKey(m, x-1, y, w, h, s); fn {
		if rn < r {
			r = rn
			f = fn
		}
	}

	if rn, fn := findKey(m, x, y-1, w, h, s); fn {
		if rn < r {
			r = rn
			f = fn
		}
	}

	if rn, fn := findKey(m, x+1, y, w, h, s); fn {
		if rn < r {
			r = rn
			f = fn
		}
	}

	if rn, fn := findKey(m, x, y+1, w, h, s); fn {
		if rn < r {
			r = rn
			f = fn
		}
	}

	return r, f
}

func countLuck(matrix []string, k int32) string {
	var m [][]rune
	var x, y int32
	for i, v := range matrix {
		var t []rune
		for j, r := range v {
			if r == 'M' {
				x = int32(i)
				y = int32(j)
			}
			t = append(t, r)
		}
		m = append(m, t)
	}

	r, _ := findKey(m, x, y, int32(len(m)), int32(len(m[0])), 0)
	if r == k {
		return "Impressed"
	}

	return "Oops!"
}

func main() {
	osFile, err := os.Open("TestCase2.txt")
	checkError(err)
	reader := bufio.NewReader(osFile)

	tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

		nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
		checkError(err)
		n := int32(nTemp)

		_, err = strconv.ParseInt(firstMultipleInput[1], 10, 64)
		checkError(err)

		var matrix []string

		for i := 0; i < int(n); i++ {
			matrixItem := readLine(reader)
			matrix = append(matrix, matrixItem)
		}

		kTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		k := int32(kTemp)

		result := countLuck(matrix, k)

		fmt.Println(result)
	}
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
