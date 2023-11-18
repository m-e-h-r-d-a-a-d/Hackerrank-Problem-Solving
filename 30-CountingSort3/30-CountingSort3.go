package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func print(result []int32, writer *bufio.Writer) {
	for _, resultItem := range result {
		fmt.Fprintf(writer, "%d", resultItem)
		fmt.Fprintf(writer, " ")

	}
	fmt.Fprintf(writer, "\n")
}

func Sort(arr []int32, writer *bufio.Writer) {
	count := make([]int32, 100)
	for _, v := range arr {
		count[v]++
	}
	sum := int32(0)
	for i := 0; i < 100; i++ {
		sum += count[i]
		count[i] = sum
	}
	print(count, writer)
}

func main() {
	osFile, _ := os.Open("TestCase0.txt")
	reader := bufio.NewReader(osFile)
	// reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	os.Setenv("OUTPUT_PATH", "output.txt")
	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")
		arrItemTemp, err := strconv.ParseInt(arrTemp[0], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	Sort(arr, writer)
	fmt.Fprintf(writer, "\n")

	writer.Flush()
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
