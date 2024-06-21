package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func weightedUniformStrings(s string, queries []int32) []string {
	output := make([]string, len(queries))
	dic := make(map[int32]struct{})

	var last rune
	counter := int32(0)
	for _, v := range s {
		if last != v {
			last = v
			counter = 0
		}

		counter += int32(v) - 96
		fmt.Println(counter)
		if _, ok := dic[counter]; !ok {
			dic[counter] = struct{}{}
		}
	}

	for i, v := range queries {
		if _, ok := dic[v]; ok {
			output[i] = "Yes"
		} else {
			output[i] = "No"
		}
	}

	return output
}

func main() {
	osFile, err := os.Open("TestCase1.txt")
	checkError(err)
	reader := bufio.NewReader(osFile)

	s := readLine(reader)
	queriesCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var queries []int32

	for i := 0; i < int(queriesCount); i++ {
		queriesItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		queriesItem := int32(queriesItemTemp)
		queries = append(queries, queriesItem)
	}

	result := weightedUniformStrings(s, queries)

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
