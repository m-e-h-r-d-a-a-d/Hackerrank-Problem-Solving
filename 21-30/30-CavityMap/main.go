package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func cavityMap(grid []string) []string {
	var output []string
	var temp [][]int
	for _, s := range grid {
		var arr []int
		for _, v := range s {
			arr = append(arr, int(v))
		}
		temp = append(temp, arr)
	}

	lenght := len(temp)
	for i := 0; i < lenght; i++ {
		s := ""
		for j := 0; j < lenght; j++ {
			v := temp[i][j]
			if i > 0 && j > 0 && i < lenght-1 && j < lenght-1 &&
				v > temp[i-1][j] && v > temp[i][j-1] && v > temp[i+1][j] && v > temp[i][j+1] {
				s = s + "X"
			} else {
				s = s + string(v)
			}
		}
		output = append(output, s)
	}
	return output
}

func main() {
	osFile, err := os.Open("TestCase1.txt")
	checkError(err)
	reader := bufio.NewReader(osFile)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	var grid []string

	for i := 0; i < int(n); i++ {
		gridItem := readLine(reader)
		grid = append(grid, gridItem)
	}

	result := cavityMap(grid)

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
