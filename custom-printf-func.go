package main

import "fmt"

func customLogf(str string, args ...interface{}) {
	fmt.Printf(str, args...)
}

func main() {
	i := 42
	customLogf("the answer is %s\n", i)
}
