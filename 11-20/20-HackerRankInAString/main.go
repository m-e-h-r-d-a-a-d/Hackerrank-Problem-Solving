package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func hackerrankInString(s string) string {
	c := []rune{'h', 'a', 'c', 'k', 'e', 'r', 'r', 'a', 'n', 'k'}
	for _, v := range s {
		if v == c[0] {
			c = c[1:]
			if len(c) == 0 {
				return "YES"
			}
		}
	}

	return "NO"
}

func main() {
	osFile, err := os.Open("TestCase1.txt")
	checkError(err)
	reader := bufio.NewReader(osFile)

	qTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	q := int32(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		s := readLine(reader)

		result := hackerrankInString(s)

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
