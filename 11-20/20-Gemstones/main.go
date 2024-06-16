package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func gemstones(arr []string) int32 {
	length := len(arr)
	count := 0
	dic := make(map[rune]int)

	for i := 0; i < int(length); i++ {
		s := arr[i]
		for _, v := range s {
			v1, ok := dic[v]
			if ok && v1 == i {
				dic[v]++
			} else if i == 0 {
				dic[v] = 1
			}
		}
	}

	for _, v := range dic {
		if v == length {
			count++
		}
	}

	return int32(count)
}

func main() {
	osFile, err := os.Open("TestCase1.txt")
	checkError(err)
	reader := bufio.NewReader(osFile)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	var arr []string

	for i := 0; i < int(n); i++ {
		arrItem := readLine(reader)
		arr = append(arr, arrItem)
	}

	result := gemstones(arr)

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
