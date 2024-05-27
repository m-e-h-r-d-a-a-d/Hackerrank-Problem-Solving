package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func toys(w []int32) int32 {
	sort.Slice(w, func(i, j int) bool {
		return w[i] < w[j]
	})
	output := int32(1)
	boundry := w[0] + int32(4)
	for i := 1; i < len(w); i++ {
		v := w[i]
		if v > boundry {
			output++
			boundry = v + int32(4)
		}
	}
	return output

}

func main() {
	osFile, err := os.Open("TestCase1.txt")
	checkError(err)
	reader := bufio.NewReader(osFile)
	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	wTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var w []int32

	for i := 0; i < int(n); i++ {
		wItemTemp, err := strconv.ParseInt(wTemp[i], 10, 64)
		checkError(err)
		wItem := int32(wItemTemp)
		w = append(w, wItem)
	}

	result := toys(w)

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
