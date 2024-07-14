package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func calculateConnected(matrix [][]int32, x int, y int) int32 {
	if matrix[x][y] == 0 {
		return 0
	}

	matrix[x][y] = 0
	count := int32(1)

	lenghX := len(matrix)
	lenghY := len(matrix[0])
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			newX := x + i
			newY := y + j
			if newX >= 0 && newY >= 0 && newX < lenghX && newY < lenghY {
				count += calculateConnected(matrix, newX, newY)
			}
		}
	}

	return count
}

func connectedCell(matrix [][]int32) int32 {
	max := int32(0)
	for i, m := range matrix {
		for j := range m {
			if matrix[i][j] == 1 {
				c := calculateConnected(matrix, i, j)
				if c > max {
					max = c
				}
			}
		}
	}

	return max
}

func main() {
	osFile, err := os.Open("TestCase1.txt")
	checkError(err)
	reader := bufio.NewReader(osFile)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	mTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	m := int32(mTemp)

	var matrix [][]int32
	for i := 0; i < int(n); i++ {
		matrixRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var matrixRow []int32
		for _, matrixRowItem := range matrixRowTemp {
			matrixItemTemp, err := strconv.ParseInt(matrixRowItem, 10, 64)
			checkError(err)
			matrixItem := int32(matrixItemTemp)
			matrixRow = append(matrixRow, matrixItem)
		}

		if len(matrixRow) != int(m) {
			panic("Bad input")
		}

		matrix = append(matrix, matrixRow)
	}

	result := connectedCell(matrix)

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
