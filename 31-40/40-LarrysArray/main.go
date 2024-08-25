package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func rotate(A []int32) {
	for A[0] > A[1] || A[0] > A[2] {
		A[0], A[1], A[2] = A[1], A[2], A[0]
	}

	return
}

func putFirst(A []int32) {
	for i := len(A) - 3; i >= 0; i-- {
		rotate(A[i : i+3])
	}
}

func larrysArray(A []int32) string {
	for i := 0; i < len(A)-2; i++ {
		putFirst(A[i:])
	}

	if A[len(A)-3] <= A[len(A)-2] && A[len(A)-2] <= A[len(A)-1] {
		return "YES"
	}

	return "NO"
}

func main() {
	osFile, _ := os.Open("TestCase2.txt")
	reader := bufio.NewReader(osFile)

	tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		n := int32(nTemp)

		ATemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

		var A []int32

		for i := 0; i < int(n); i++ {
			AItemTemp, err := strconv.ParseInt(ATemp[i], 10, 64)
			checkError(err)
			AItem := int32(AItemTemp)
			A = append(A, AItem)
		}

		result := larrysArray(A)

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
