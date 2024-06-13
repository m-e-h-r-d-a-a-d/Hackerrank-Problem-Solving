package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func anagram(s string) int32 {
	length := len(s)
	if length%2 == 1 {
		return -1
	}

	s1 := s[:length/2]
	s2 := s[length/2:]

	dic := make(map[rune]int32)
	for _, v := range s1 {
		if _, ok := dic[v]; ok {
			dic[v]++
		} else {
			dic[v] = 1
		}
	}

	for _, v := range s2 {
		if _, ok := dic[v]; ok {
			dic[v]--
		} else {
			dic[v] = -1
		}
	}

	output := int32(0)
	for _, v := range dic {
		if v > 0 {
			output += v
		} else {
			output -= v
		}
	}

	return output / 2
}

func main() {
	osFile, err := os.Open("TestCase1.txt")
	checkError(err)
	reader := bufio.NewReader(osFile)

	qTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	q := int32(qTemp)

	for i := 0; i < int(q); i++ {
		s := readLine(reader)
		result := anagram(s)
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
