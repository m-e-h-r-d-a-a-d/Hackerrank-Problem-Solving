package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func bobAndBen(trees [][]int32) string {
	xor := int32(0)
	for _, t := range trees {
		if t[0] == 0 || t[0] == 2 {
			xor ^= 0
		} else {
			xor ^= (t[0]-1)%2 + 1
		}
	}
	if xor == 0 {
		return "BEN"
	}
	return "BOB"
}

func main() {
	osFile, err := os.Open("TestCase1.txt")
	checkError(err)
	reader := bufio.NewReader(osFile)

	gTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	g := int32(gTemp)

	for gItr := 0; gItr < int(g); gItr++ {
		nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		n := int32(nTemp)

		var trees [][]int32
		for i := 0; i < int(n); i++ {
			treesRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

			var treesRow []int32
			for _, treesRowItem := range treesRowTemp {
				treesItemTemp, err := strconv.ParseInt(treesRowItem, 10, 64)
				checkError(err)
				treesItem := int32(treesItemTemp)
				treesRow = append(treesRow, treesItem)
			}

			if len(treesRow) != 2 {
				panic("Bad input")
			}

			trees = append(trees, treesRow)
		}

		result := bobAndBen(trees)

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
