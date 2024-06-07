package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func happyLadybugs(b string) string {
	hashMap := make(map[rune]bool)
	isUnderScore := false

	for _, v := range b {
		if v == '_' {
			isUnderScore = true
			continue
		}
		if _, ok := hashMap[v]; ok {
			hashMap[v] = true
		} else {
			hashMap[v] = false
		}
	}

	if !isUnderScore {
		last := 's'
		counter := 0
		for _, v := range b {
			if v == last {
				counter++
				continue
			}

			if counter == 1 {
				return "NO"
			}

			counter = 1
			last = v

		}

		if counter == 1 {
			return "NO"
		}

		return "YES"
	}

	for _, v := range hashMap {
		if !v {
			return "NO"
		}
	}

	return "YES"
}

func main() {
	osFile, err := os.Open("TestCase1.txt")
	checkError(err)
	reader := bufio.NewReader(osFile)

	gTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	g := int32(gTemp)

	for gItr := 0; gItr < int(g); gItr++ {
		_, err = strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)

		b := readLine(reader)

		result := happyLadybugs(b)

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
