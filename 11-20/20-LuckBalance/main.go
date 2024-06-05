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

func luckBalance(k int32, contests [][]int32) int32 {
	var importantCount, luck int32
	for i := 0; i < len(contests); i++ {
		if contests[i][1] == 1 {
			importantCount++
		}
	}

	winCount := importantCount - k
	sort.Slice(contests, func(i int, j int) bool {
		if contests[i][1] > contests[j][1] {
			return true
		} else if contests[i][1] < contests[j][1] {
			return false
		}

		return contests[i][0] < contests[j][0]
	})

	fmt.Println(contests)
	for i := 0; i < len(contests); i++ {
		if winCount > 0 {
			luck -= contests[i][0]
			winCount--
		} else {
			luck += contests[i][0]
		}
	}
	return luck
}

func main() {
	osFile, err := os.Open("TestCase1.txt")
	checkError(err)
	reader := bufio.NewReader(osFile)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	kTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	k := int32(kTemp)

	var contests [][]int32
	for i := 0; i < int(n); i++ {
		contestsRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var contestsRow []int32
		for _, contestsRowItem := range contestsRowTemp {
			contestsItemTemp, err := strconv.ParseInt(contestsRowItem, 10, 64)
			checkError(err)
			contestsItem := int32(contestsItemTemp)
			contestsRow = append(contestsRow, contestsItem)
		}

		if len(contestsRow) != 2 {
			panic("Bad input")
		}

		contests = append(contests, contestsRow)
	}

	result := luckBalance(k, contests)

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
