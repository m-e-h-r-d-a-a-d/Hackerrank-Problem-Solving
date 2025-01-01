package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func IsBeautiful(a string, s string) bool {
	if len(s) == 0 {
		return true
	}

	if len(a) == 0 {
		return false
	}

	if s[0] == '0' {
		return false
	}

	if len(a) > len(s) {
		return false
	}

	if a != s[:len(a)] {
		return false
	}

	aInt, err := strconv.Atoi(a)
	if err != nil {
		return false
	}

	return IsBeautiful(strconv.Itoa(aInt+1), s[len(a):])
}

func separateNumbers(s string) {
	for i := 1; i <= len(s)/2; i++ {
		if IsBeautiful(s[:i], s) {
			fmt.Println("YES", s[:i])
			return
		}
	}

	fmt.Println("NO")
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

		separateNumbers(s)
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
