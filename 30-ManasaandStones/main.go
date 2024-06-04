package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func stones(n int32, a int32, b int32) []int32 {
	var output []int32
	if a > b {
		a, b = b, a
	}
	n--
	var last int32
	for i := int32(0); i <= n; i++ {
		v := i*b + (n-i)*a
		if last != v {
			output = append(output, v)
		}
		last = v
	}

	return output
}

func main() {
	osFile, err := os.Open("TestCase2.txt")
	checkError(err)
	reader := bufio.NewReader(osFile)

	TTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	T := int32(TTemp)

	for TItr := 0; TItr < int(T); TItr++ {
		nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		n := int32(nTemp)

		aTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		a := int32(aTemp)

		bTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		b := int32(bTemp)

		result := stones(n, a, b)

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
