package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func makingAnagrams(s1 string, s2 string) int32 {
	var output int32
	dic := make(map[rune]int32)
	for _, v := range s1 {
		if _, ok := dic[v]; ok {
			dic[v] = dic[v] + 1
		} else {
			dic[v] = 1
		}
	}

	for _, v := range s2 {
		if _, ok := dic[v]; ok {
			dic[v] = dic[v] - 1
		} else {
			dic[v] = -1
		}
	}

	for _, v := range dic {
		if v < 0 {
			output -= v
		} else {
			output += v
		}

	}

	return output
}

func main() {
	osFile, err := os.Open("TestCase1.txt")
	checkError(err)
	reader := bufio.NewReader(osFile)

	s1 := readLine(reader)
	s2 := readLine(reader)

	result := makingAnagrams(s1, s2)
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
