package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func calculateNimSum(p []int32) int32 {
	n := int32(0)
	for _, v := range p {
		n ^= v
	}
	return n
}

func pokerNim(c []int32) string {
	n := calculateNimSum(c)

	if n == 0 {
		return "Second"
	}

	return "First"
}

func main() {
	osFile, err := os.Open("TestCase1.txt")
	checkError(err)
	reader := bufio.NewReader(osFile)

	tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

		nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
		checkError(err)
		n := int32(nTemp)

		_, err = strconv.ParseInt(firstMultipleInput[1], 10, 64)
		checkError(err)

		cTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

		var c []int32

		for i := 0; i < int(n); i++ {
			cItemTemp, err := strconv.ParseInt(cTemp[i], 10, 64)
			checkError(err)
			cItem := int32(cItemTemp)
			c = append(c, cItem)
		}

		result := pokerNim(c)

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
