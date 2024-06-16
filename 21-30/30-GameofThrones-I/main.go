package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func gameOfThrones(s string) string {
	dic := make(map[rune]bool)
	for _, v := range s {
		if c, ok := dic[v]; !ok || !c {
			dic[v] = true
		} else {
			dic[v] = false
		}
	}

	count := 0
	for _, v := range dic {
		if v {
			count++
			if count > 1 {
				return "NO"
			}
		}
	}

	return "YES"
}

func main() {
	osFile, err := os.Open("TestCase1.txt")
	checkError(err)
	reader := bufio.NewReader(osFile)

	s := readLine(reader)

	result := gameOfThrones(s)

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
