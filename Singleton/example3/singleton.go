package main

import (
	"fmt"
)

var (
	singletonInstance *singleton
)

type singleton struct{}

// wrong example
func GetInstance() *singleton {
	singletonInstance = &singleton{}
	fmt.Println(&singletonInstance)
	return singletonInstance
}

// wrong example
func GetInstance2() *singleton {
	singletonInstance2 := &singleton{}
	fmt.Println(&singletonInstance2)
	return singletonInstance2
}
