package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func isPalindrome(s string) (bool, int) {
	lenght := len(s)
	for i := 0; i <= lenght/2; i++ {
		if s[i] != s[lenght-i-1] {
			return false, i
		}
	}

	return true, -1
}

func palindromeIndex(s string) int32 {
	p, idx1 := isPalindrome(s)
	if p {
		return -1
	}

	tempS := ""
	if idx1 == 0 {
		tempS = s[idx1+1:]
	} else {
		tempS = s[:idx1] + s[idx1+1:]
	}

	p, _ = isPalindrome(tempS)

	if p {
		return int32(idx1)
	} else {
		return int32(len(s) - idx1 - 1)
	}

}

func main() {
	osFile, err := os.Open("TestCase2.txt")
	checkError(err)
	reader := bufio.NewReader(osFile)

	qTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	q := int32(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		s := readLine(reader)

		result := palindromeIndex(s)

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
