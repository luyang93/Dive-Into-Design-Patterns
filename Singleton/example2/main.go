package main

import "fmt"

func main() {
	for i := 0; i < 1000; i++ {
		go GetInstance()
	}
	fmt.Scanln()
}
