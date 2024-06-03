package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func beautifulBinaryString(b string) int32 {
	output := int32(0)
	var temp []string
	temp = append(temp, string(b[0]))
	temp = append(temp, string(b[1]))
	for i := 2; i < len(b); i++ {
		v := string(b[i])
		check := temp[i-2] + temp[i-1] + v
		if check == "010" {
			output++
			temp = append(temp, "1")
		} else {
			temp = append(temp, v)
		}
	}

	return output
}

func main() {
	osFile, err := os.Open("TestCase1.txt")
	checkError(err)
	reader := bufio.NewReader(osFile)

	_, err = strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	b := readLine(reader)

	result := beautifulBinaryString(b)
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
