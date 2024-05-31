package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func funnyString(s string) string {
	var asci []int
	for i := 1; i < len(s); i++ {
		a := int(s[i]) - int(s[i-1])
		if a < 0 {
			a *= -1
		}
		asci = append(asci, a)
	}

	fmt.Println(s)
	fmt.Println(asci)
	length := len(asci)

	for i := 0; i < length/2; i++ {
		if asci[length-i-1]-asci[i] != 0 {
			return "Not Funny"
		}
	}

	return "Funny"
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

		result := funnyString(s)

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
