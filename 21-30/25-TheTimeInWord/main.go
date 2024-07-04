package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var (
	clocks = []string{
		"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
		"ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen",
		"twenty", "twenty one", "twenty two", "twenty three", "twenty four", "twenty five", "twenty six", "twenty seven", "twenty eight", "twenty nine",
	}

	tensclocks = []string{
		"", "", "", "thirty", "forty", "fifty",
	}
)

func timeInWords(h int32, m int32) string {

	if m == 0 {
		return clocks[h] + " o' clock"
	}

	if m == 15 {
		return "quarter past " + clocks[h]
	}

	if m == 30 {
		return "half past " + clocks[h]
	}

	if m == 45 {
		return "quarter to " + clocks[h+1]
	}

	if m < 30 {
		return clocks[m] + " " + pluralize(m, "minute") + " past " + clocks[h]
	}

	return clocks[60-m] + " " + pluralize(int32(60)-m, "minute") + " to " + clocks[h+1]
}

func pluralize(n int32, unit string) string {
	if n == 1 {
		return unit
	}
	return unit + "s"
}

func main() {
	osFile, err := os.Open("TestCase2.txt")
	checkError(err)
	reader := bufio.NewReader(osFile)

	hTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	h := int32(hTemp)

	mTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	m := int32(mTemp)

	result := timeInWords(h, m)

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
