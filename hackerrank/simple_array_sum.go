package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the simpleArraySum function below.
 */
func simpleArraySum(ar []int32) int32 {
	var sum int32 = 0
	if len(ar) == 1 {
		return ar[0]
	}
	for i := 0; i < len(ar); i++ {
		sum += ar[i]
	}

	return sum
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	arCount, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	arTemp := strings.Split(readLine(reader), " ")

	var ar []int32

	for arItr := 0; arItr < int(arCount); arItr++ {
		arItemTemp, err := strconv.ParseInt(arTemp[arItr], 10, 64)
		checkError(err)
		arItem := int32(arItemTemp)
		ar = append(ar, arItem)
	}

	result := simpleArraySum(ar)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}
