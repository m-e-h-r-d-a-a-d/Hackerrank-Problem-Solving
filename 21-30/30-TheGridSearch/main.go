package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func remove(slice []string, i int) []string {
	return append(slice[:i], slice[i+1:]...)
}

// func gridSearch(G []string, P []string) string {
// 	for i, p := range P {
// 		lenght := len(p)
// 		isFound := false

// 	out:
// 		for i, g := range G {
// 			isFound := false

// 			for i := 0; i <= len(g)-lenght; i++ {
// 				if p == g[i:lenght+i] {
// 					isFound = true
// 					break out
// 				}
// 			}
// 		}

// 		if isFound == false {
// 			return "NO"
// 		}
// 	}

// 	return "YES"
// }

func gridSearch(G []string, P []string) string {
	lenG := len(G[0])
	lenP := len(P[0])
	for i := 0; i <= len(G)-len(P); i++ {
		for j := 0; j <= lenG-lenP; j++ {
			isEqual := true
			for k := 0; k < len(P); k++ {
				if G[i+k][j:j+lenP] != P[k] {
					isEqual = false
					break
				}
			}

			if isEqual {
				return "YES"
			}
		}
	}
	return "NO"
}

func main() {
	osFile, err := os.Open("TestCase1.txt")
	checkError(err)
	reader := bufio.NewReader(osFile)
	tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	t := int32(tTemp)
	for tItr := 0; tItr < int(t); tItr++ {
		firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

		RTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
		checkError(err)
		R := int32(RTemp)

		_, err = strconv.ParseInt(firstMultipleInput[1], 10, 64)
		checkError(err)

		var G []string

		for i := 0; i < int(R); i++ {
			GItem := readLine(reader)
			G = append(G, GItem)
		}

		secondMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

		rTemp, err := strconv.ParseInt(secondMultipleInput[0], 10, 64)
		checkError(err)
		r := int32(rTemp)

		_, err = strconv.ParseInt(secondMultipleInput[1], 10, 64)
		checkError(err)

		var P []string

		for i := 0; i < int(r); i++ {
			PItem := readLine(reader)
			P = append(P, PItem)
		}

		result := gridSearch(G, P)

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
