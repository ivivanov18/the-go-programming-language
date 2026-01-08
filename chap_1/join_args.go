package main

import (
	"os"
	"fmt"
	"strings"
	"time"
)

func main() {
	arr := generateArray(100000, "hi")

	start := time.Now()
	joinConcat(arr)
	secs := time.Since(start).Seconds()
	fmt.Printf("It took %f seconds to run manual string concatenation\n", secs)

	start = time.Now()
	joinUtility(arr)
	secs = time.Since(start).Seconds()
	fmt.Printf("It took %f seconds to run string concatenation using strings.Join\n", secs)
}

func joinConcat(args []string) string {
	s, sep := "", ""

	for _, arg := range args{
		s += sep + arg
		sep = " "
	}
	return s
}

func joinUtility(args []string) string {
	return strings.Join(args, " ")
}

func echoCommandInvoked() {
	fmt.Printf("Invoking %s\n", os.Args[0])
}

func echoRemainingArguments() {
	for i, arg := range os.Args[1:] {
		fmt.Printf("arg at pos %d is %s\n", i, arg)
	}
}

func generateArray(nbElts int, value string) []string {
	arr := make([]string, nbElts)

	for i := range arr {
		arr[i] = value
	}
	return arr
}
