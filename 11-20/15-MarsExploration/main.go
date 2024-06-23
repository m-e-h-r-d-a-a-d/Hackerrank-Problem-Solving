package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func marsExploration(s string) int32 {
	sos := "SOS"
	j := 0
	output := int32(0)
	for i := 0; i < len(s); i++ {
		if s[i] != sos[j] {
			output++
		}

		j++
		if j == 3 {
			j = 0
		}
	}
	return output
}

func main() {
	osFile, err := os.Open("TestCase1.txt")
	checkError(err)
	reader := bufio.NewReader(osFile)

	s := readLine(reader)

	result := marsExploration(s)
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
