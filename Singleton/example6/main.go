package main

import "fmt"

func main() {
	instance1 := GetInstance(0)
	instance2 := GetInstance(0)

	fmt.Println(instance1 == instance2) // Output: true

	instance3 := GetInstance(2)
	instance4 := GetInstance(2)
	instance5 := GetInstance(2)

	fmt.Println(instance3 == instance4) // Output: false
	fmt.Println(instance4 == instance5) // Output: false
	fmt.Println(instance3 == instance5) // Output: true

}
