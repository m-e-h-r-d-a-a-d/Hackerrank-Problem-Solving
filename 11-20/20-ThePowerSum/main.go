package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func powerSumRecursive(X, N, sum, begin int32) int32 {
	if sum == X {
		return 1
	}
	if sum > X || begin > int32(math.Pow(float64(X), 1.0/float64(N))) {
		return 0
	}

	return powerSumRecursive(X, N, sum, begin+1) + powerSumRecursive(X, N, sum+int32(math.Pow(float64(begin), float64(N))), begin+1)
}

func powerSum(X int32, N int32) int32 {
	return powerSumRecursive(X, N, int32(0), int32(1))
}

func main() {
	osFile, err := os.Open("TestCase1.txt")
	checkError(err)
	reader := bufio.NewReader(osFile)

	XTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	X := int32(XTemp)

	NTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	N := int32(NTemp)

	result := powerSum(X, N)

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
