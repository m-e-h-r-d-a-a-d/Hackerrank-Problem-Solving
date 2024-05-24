package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

/*
 * Complete the 'pangrams' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func pangrams(s string) string {
	// Write your code here
	counter := 0
	a := int('a')
	z := int('z')
	dic := make(map[int](bool))
	for _, v := range strings.ToLower(s) {
		c := int(v)
		if c >= a && c <= z {
			_, ok := dic[c]
			if !ok {
				dic[c] = true
				counter++
				if counter == 26 {
					fmt.Println(counter)
					return "pangram"
				}
			}

		}
	}
	return "not pangram"
}

func main() {
	osFile, _ := os.Open("TestCase1.txt")
	reader := bufio.NewReader(osFile)
	s := readLine(reader)

	result := pangrams(s)

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
