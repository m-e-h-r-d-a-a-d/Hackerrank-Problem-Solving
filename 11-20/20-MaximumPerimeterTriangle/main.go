package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func canFormTriangle(a, b, c int32) bool {
	return (a+b > c) && (a+c > b) && (b+c > a)
}

func maximumPerimeterTriangle(sticks []int32) []int32 {
	sort.Slice(sticks, func(i, j int) bool {
		return sticks[i] < sticks[j]
	})

	for i := len(sticks) - 1; i > 1; i-- {
		if canFormTriangle(sticks[i-2], sticks[i-1], sticks[i]) {
			return []int32{sticks[i-2], sticks[i-1], sticks[i]}
		}
	}

	return []int32{-1}
}

func main() {
	osFile, err := os.Open("TestCase1.txt")
	checkError(err)
	reader := bufio.NewReader(osFile)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	sticksTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var sticks []int32

	for i := 0; i < int(n); i++ {
		sticksItemTemp, err := strconv.ParseInt(sticksTemp[i], 10, 64)
		checkError(err)
		sticksItem := int32(sticksItemTemp)
		sticks = append(sticks, sticksItem)
	}

	result := maximumPerimeterTriangle(sticks)

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
