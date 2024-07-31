package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func recursive(s string) string {
	lenght := len(s)
	for i := 1; i < lenght; i++ {
		if s[i] == s[i-1] {
			startIndex := i - 1
			endIndex := i
			if startIndex >= 0 && endIndex < lenght {
				return recursive(s[:startIndex] + s[endIndex+1:])
			} else {
				if startIndex == 0 {
					return recursive(s[endIndex+1:])
				}

				if endIndex == lenght {
					return recursive(s[:startIndex])
				}

				return ""
			}
		}
	}
	return s
}

func superReducedString(s string) string {
	output := recursive(s)
	if output == "" {
		return "Empty String"
	}

	return output
}

func main() {
	f, err := os.Open("TestCase1.txt")
	checkError(err)
	reader := bufio.NewReader(f)

	s := readLine(reader)

	result := superReducedString(s)

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
