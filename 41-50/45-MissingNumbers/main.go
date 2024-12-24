package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
)

func missingNumbers(arr []int32, brr []int32) []int32 {
	dic1 := make(map[int32]int)
	dic2 := make(map[int32]int)
	output := []int32{}
	for _, v := range arr {
		if _, ok := dic1[v]; ok {
			dic1[v]++
		} else {
			dic1[v] = 1
		}
	}

	for _, v := range brr {
		if _, ok := dic2[v]; ok {
			dic2[v]++
		} else {
			dic2[v] = 1
		}
	}

	for _, v := range brr {
		if _, ok := dic2[v]; !ok {
			continue
		}

		_, ok := dic1[v]

		if !ok || dic2[v] != dic1[v] {
			output = append(output, v)
			delete(dic2, v)
		}
	}

	slices.Sort(output)
	return output
}

func main() {
	osFile, err := os.Open("TestCase2.txt")
	checkError(err)
	reader := bufio.NewReader(osFile)

	defer osFile.Close()

	writer := bufio.NewWriterSize(osFile, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	mTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	m := int32(mTemp)

	brrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var brr []int32

	for i := 0; i < int(m); i++ {
		brrItemTemp, err := strconv.ParseInt(brrTemp[i], 10, 64)
		checkError(err)
		brrItem := int32(brrItemTemp)
		brr = append(brr, brrItem)
	}

	result := missingNumbers(arr, brr)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%d", resultItem)

		if i != len(result)-1 {
			fmt.Fprintf(writer, " ")
		}
	}

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
