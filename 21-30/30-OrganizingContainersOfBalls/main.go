package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func organizingContainers(container [][]int32) string {
	keys := make(map[int32]struct{})
	lengthX := len(container)
	lengthY := len(container[0])
	for i := 0; i < lengthX; i++ {
		sum := int32(0)
		for j := 0; j < lengthY; j++ {
			sum += container[i][j]
		}
		keys[sum] = struct{}{}
	}

	for j := 0; j < lengthY; j++ {
		sum := int32(0)
		for i := 0; i < lengthX; i++ {
			sum += container[i][j]
		}
		if _, ok := keys[sum]; !ok {
			return "Impossible"
		}
	}

	return "Possible"
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

		var container [][]int32
		for i := 0; i < int(n); i++ {
			containerRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

			var containerRow []int32
			for _, containerRowItem := range containerRowTemp {
				containerItemTemp, err := strconv.ParseInt(containerRowItem, 10, 64)
				checkError(err)
				containerItem := int32(containerItemTemp)
				containerRow = append(containerRow, containerItem)
			}

			if len(containerRow) != int(n) {
				panic("Bad input")
			}

			container = append(container, containerRow)
		}

		result := organizingContainers(container)

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
