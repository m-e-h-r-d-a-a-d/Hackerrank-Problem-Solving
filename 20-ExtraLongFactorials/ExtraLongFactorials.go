package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'extraLongFactorials' function below.
 *
 * The function accepts INTEGER n as parameter.
 */

func multi(input1 []int, x int) []int {

	s := strconv.Itoa(x)
	lenght := len(s)
	input2 := make([]int, lenght)
	result := make([]int, lenght+len(input1))
	for i, v := range s {
		value, _ := strconv.Atoi(string(v))
		input2[lenght-i-1] = value
	}

	for i := 0; i < len(input1); i++ {
		temp := 0
		var tempInput []int
		for k := 0; k < i; k++ {
			tempInput = append(tempInput, 0)
		}

		for j := 0; j < len(input2); j++ {
			v := input1[i]*input2[j] + temp
			if v > 9 {
				temp = v / 10
				v = v % 10
			} else {
				temp = 0
			}
			tempInput = append(tempInput, v)
		}

		if temp > 0 {
			tempInput = append(tempInput, temp)
		}

		temp = 0
		index := 0
		for index = 0; index < len(tempInput); index++ {
			v := result[index] + tempInput[index] + temp
			if v > 9 {
				temp = v / 10
				v = v % 10
			} else {
				temp = 0
			}
			result[index] = v
		}

		for temp > 0 {
			v := result[index] + temp
			if v > 9 {
				temp = v / 10
				v = v % 10
			} else {
				temp = 0
			}
			result[index] = v
			index++
		}

	}
	index := len(result)
	for i := len(result) - 1; i >= 0; i-- {
		if result[i] != 0 {
			break
		}
		index = i
	}

	return result[:index]
}

func extraLongFactorials(n int32) {
	// Write your code here
	result := []int{1}
	for i := 2; i <= int(n); i++ {
		result = multi(result, i)
	}
	for i := len(result) - 1; i >= 0; i-- {
		fmt.Print(result[i])
	}
}

func main() {
	osFile, _ := os.Open("TestCase1.txt")
	reader := bufio.NewReader(osFile)
	// reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	extraLongFactorials(n)
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
