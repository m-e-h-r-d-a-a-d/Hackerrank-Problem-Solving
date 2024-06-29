package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func fairRations(B []int32) string {
	count := 0
	length := len(B) - 1
	for i, v := range B {
		if int(v)%2 == 1 {
			if i == length {
				return "NO"
			}
			count += 2
			B[i+1]++
		}
	}

	return strconv.Itoa(count)
}

func main() {
	osFile, err := os.Open("TestCase1.txt")
	checkError(err)
	reader := bufio.NewReader(osFile)

	NTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	N := int32(NTemp)

	BTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var B []int32

	for i := 0; i < int(N); i++ {
		BItemTemp, err := strconv.ParseInt(BTemp[i], 10, 64)
		checkError(err)
		BItem := int32(BItemTemp)
		B = append(B, BItem)
	}

	result := fairRations(B)
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
