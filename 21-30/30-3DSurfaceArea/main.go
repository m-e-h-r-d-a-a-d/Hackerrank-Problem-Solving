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
 * Complete the 'surfaceArea' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts 2D_INTEGER_ARRAY A as parameter.
 */

func surfaceArea(A [][]int32) int32 {
	w := len(A)
	h := len(A[0])
	lrSide := int32(0)
	frSide := int32(0)
	udSide := int32(0)

	for i := 0; i < w; i++ {
		last := int32(0)
		for j := 0; j < h; j++ {
			if A[i][j] > 0 {
				udSide++
			}

			v := A[i][j] - last
			if v > 0 {
				lrSide += v
			}

			last = A[i][j]
		}
	}

	for i := 0; i < h; i++ {
		last := int32(0)
		for j := 0; j < w; j++ {
			v := A[j][i] - last
			if v > 0 {
				frSide += v
			}

			last = A[j][i]
		}
	}

	return (lrSide + frSide + udSide) * 2
}

func main() {
	osFile, err := os.Open("TestCase1.txt")
	checkError(err)
	reader := bufio.NewReader(osFile)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	HTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	H := int32(HTemp)

	WTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	W := int32(WTemp)

	var A [][]int32
	for i := 0; i < int(H); i++ {
		ARowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var ARow []int32
		for _, ARowItem := range ARowTemp {
			AItemTemp, err := strconv.ParseInt(ARowItem, 10, 64)
			checkError(err)
			AItem := int32(AItemTemp)
			ARow = append(ARow, AItem)
		}

		if len(ARow) != int(W) {
			panic("Bad input")
		}

		A = append(A, ARow)
	}

	result := surfaceArea(A)

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
