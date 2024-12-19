package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func decentNumber(n int32) {
	intN := int(n)
	count := 0
	decent := false
	for i := 0; i <= intN; i++ {
		if i%5 == 0 && (intN-i)%3 == 0 {
			count = i
			decent = true
			break
		}
	}

	if !decent {
		fmt.Println(-1)
		return
	}

	for i := 0; i < intN-count; i++ {
		fmt.Print("5")
	}

	for i := 0; i < count; i++ {
		fmt.Print("3")
	}

	fmt.Println()
}

func main() {
	osFile, err := os.Open("TestCase1.txt")
	checkError(err)
	reader := bufio.NewReader(osFile)

	tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		n := int32(nTemp)

		decentNumber(n)
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
