package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func stockmax(prices []int32) int64 {
	if len(prices) == 0 {
		return 0
	}
	profit := int64(0)
	maxPrice := prices[len(prices)-1]
	for i := len(prices) - 2; i >= 0; i-- {
		if prices[i] < maxPrice {
			profit += int64(maxPrice - prices[i])
		} else {
			maxPrice = prices[i]
		}
	}
	return profit
}

func main() {
	osFile, err := os.Open("TestCase1.txt")
	checkError(err)
	reader := bufio.NewReader(osFile)
	tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		n := int32(nTemp)

		pricesTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

		var prices []int32

		for i := 0; i < int(n); i++ {
			pricesItemTemp, err := strconv.ParseInt(pricesTemp[i], 10, 64)
			checkError(err)
			pricesItem := int32(pricesItemTemp)
			prices = append(prices, pricesItem)
		}

		result := stockmax(prices)

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
