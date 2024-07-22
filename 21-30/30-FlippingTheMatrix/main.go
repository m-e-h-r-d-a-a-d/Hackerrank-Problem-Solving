package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func flippingMatrix(matrix [][]int32) int32 {
	length := len(matrix) / 2
	sum := int32(0)
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			m1 := max(matrix[i][j], matrix[i][2*length-j-1])
			m2 := max(matrix[2*length-i-1][j], matrix[2*length-i-1][2*length-j-1])
			sum += max(m1, m2)
		}
	}
	return sum
}

func max(a, b int32) int32 {
	if a > b {
		return a
	}
	return b
}

func main() {
	osFile, err := os.Open("TestCase1.txt")
	checkError(err)
	reader := bufio.NewReader(osFile)
	qTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	q := int32(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		n := int32(nTemp)

		var matrix [][]int32
		for i := 0; i < 2*int(n); i++ {
			matrixRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

			var matrixRow []int32
			for _, matrixRowItem := range matrixRowTemp {
				matrixItemTemp, err := strconv.ParseInt(matrixRowItem, 10, 64)
				checkError(err)
				matrixItem := int32(matrixItemTemp)
				matrixRow = append(matrixRow, matrixItem)
			}

			if len(matrixRow) != 2*int(n) {
				panic("Bad input")
			}

			matrix = append(matrix, matrixRow)
		}

		result := flippingMatrix(matrix)

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
