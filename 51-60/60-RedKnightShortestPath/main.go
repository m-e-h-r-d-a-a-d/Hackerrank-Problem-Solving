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

func printShortestPath(n int32, i_start int32, j_start int32, i_end int32, j_end int32) {
	if (i_start-i_end)%2 != 0 {
		fmt.Println("Impossible")
		return
	}

	path := make([]int, 0)
	counter := 0
	for i_start != i_end {
		if i_start > i_end {
			if j_start > j_end {
				path = append(path, 1)
				j_start--
				i_start -= 2
			} else {
				path = append(path, 2)
				j_start++
				i_start -= 2
			}
		} else {
			if j_start > j_end {
				path = append(path, 5)
				j_start--
				i_start += 2
			} else {
				path = append(path, 4)
				j_start++
				i_start += 2
			}
		}
		counter++
	}

	if (j_start-j_end)%2 != 0 {
		fmt.Println("Impossible")
		return
	}

	for j_start != j_end {
		if j_start > j_end {
			path = append(path, 6)
			j_start -= 2
		} else {
			path = append(path, 3)
			j_start += 2
		}
		counter++
	}

	sort.Slice(path, func(i, j int) bool {
		return path[i] < path[j]
	})

	output := ""
	for i := 0; i < len(path)-1; i++ {
		output += convert(path[i]) + " "
	}

	output += convert(path[len(path)-1])

	fmt.Println(counter)
	fmt.Println(output)
}

func convert(i int) string {
	switch i {
	case 1:
		return "UL"
	case 2:
		return "UR"
	case 3:
		return "R"
	case 4:
		return "LR"
	case 5:
		return "LL"
	case 6:
		return "L"
	}

	return ""
}

func main() {
	f, err := os.Open("TestCase2.txt")
	checkError(err)
	reader := bufio.NewReader(f)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	i_startTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	i_start := int32(i_startTemp)

	j_startTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	j_start := int32(j_startTemp)

	i_endTemp, err := strconv.ParseInt(firstMultipleInput[2], 10, 64)
	checkError(err)
	i_end := int32(i_endTemp)

	j_endTemp, err := strconv.ParseInt(firstMultipleInput[3], 10, 64)
	checkError(err)
	j_end := int32(j_endTemp)

	printShortestPath(n, i_start, j_start, i_end, j_end)
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
