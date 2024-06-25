package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func isAlternating(s string, a, b byte) bool {
	prev := byte(0)
	for i := 0; i < len(s); i++ {
		if s[i] == a || s[i] == b {
			if s[i] == prev {
				return false
			}
			prev = s[i]
		}
	}
	return true
}

func alternate(s string) int {
	uniqueChars := make(map[byte]bool)
	for i := 0; i < len(s); i++ {
		uniqueChars[s[i]] = true
	}

	maxLength := 0
	for a := range uniqueChars {
		for b := range uniqueChars {
			if a != b {
				if isAlternating(s, a, b) {
					length := 0
					for i := 0; i < len(s); i++ {
						if s[i] == a || s[i] == b {
							length++
						}
					}
					maxLength = max(maxLength, length)
				}
			}
		}
	}

	return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	osFile, err := os.Open("TestCase1.txt")
	checkError(err)
	reader := bufio.NewReader(osFile)

	readLine(reader)
	checkError(err)

	s := readLine(reader)

	result := alternate(s)

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
