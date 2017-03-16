package main

import "fmt"

func main() {
	rate := 42

	// this condition can never be true
	if rate > 60 && rate < 40 {
		fmt.Println("rate %:", rate)
	}
}
