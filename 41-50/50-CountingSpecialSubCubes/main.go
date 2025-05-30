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
 * Complete the 'specialSubCubes' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts INTEGER_ARRAY cube as parameter.
 */

func specialSubCubes(cube []int32) []int32 {
	fmt.Println(cube)
	return []int32{}
}

func main() {
	osFile, err := os.Open("TestCase1.txt")
	checkError(err)
	reader := bufio.NewReader(osFile)

	qTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	q := int32(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		n := int32(nTemp)

		cubeTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

		var cube []int32

		for i := 0; i < int(n)*int(n)*int(n); i++ {
			cubeItemTemp, err := strconv.ParseInt(cubeTemp[i], 10, 64)
			checkError(err)
			cubeItem := int32(cubeItemTemp)
			cube = append(cube, cubeItem)
		}

		result := specialSubCubes(cube)

		for _, resultItem := range result {
			fmt.Println(resultItem)

		}

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
