package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strings"
)

func encryption(s string) string {
	output := ""
	length := len(s)
	fmt.Println(length)
	rows := int(math.Sqrt(float64(len(s))))
	cols := rows
	if length > rows*cols {
		cols++
	}

	for i := 0; i < cols; i++ {
		for j := 0; j <= rows; j++ {
			idx := j*cols + i
			if idx >= length {
				continue
			}

			output += string(s[idx])
		}
		output += " "
	}

	return output
}

func main() {
	osFile, err := os.Open("TestCase1.txt")
	checkError(err)
	reader := bufio.NewReader(osFile)

	t := strings.TrimSpace(readLine(reader))
	checkError(err)

	result := encryption(t)

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
