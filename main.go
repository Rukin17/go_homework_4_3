package main

import "fmt"

func main() {
	counter := func() func() int {
		count := 0
		return func() int {
			count++
			return count
		}
	}()

	for range 5 {
		fmt.Println(counter())
	}
}
