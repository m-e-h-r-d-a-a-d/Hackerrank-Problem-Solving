package main

import (
	"bufio"
	"fmt"
	"io"
	"math/big"
	"os"
	"strconv"
	"strings"
)

func fibonacciModified(t1 *big.Int, t2 *big.Int, n int32) *big.Int {
	var new big.Int
	for i := 3; i <= int(n); i++ {
		new.Mul(t2, t2).Add(&new, t1)
		t1.Set(t2)
		t2.Set(&new)
	}

	return t2
}

func main() {
	osFile, err := os.Open("TestCase1.txt")
	checkError(err)
	reader := bufio.NewReader(osFile)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	t1Temp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	t1 := big.NewInt(int64(t1Temp))

	t2Temp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	t2 := big.NewInt(int64(t2Temp))

	nTemp, err := strconv.ParseInt(firstMultipleInput[2], 10, 64)
	checkError(err)
	n := int32(nTemp)

	result := fibonacciModified(t1, t2, n)

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
