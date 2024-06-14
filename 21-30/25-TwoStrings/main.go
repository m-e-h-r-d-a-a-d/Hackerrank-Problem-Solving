package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func twoStrings(s1 string, s2 string) string {
	dic := make(map[rune]struct{})

	for _, v := range s1 {
		dic[v] = struct{}{}
	}

	for _, v := range s2 {
		if _, ok := dic[v]; ok {
			return "YES"
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

	for i := 0; i < int(qTemp); i++ {
		s1 := readLine(reader)
		s2 := readLine(reader)
		result := twoStrings(s1, s2)
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
