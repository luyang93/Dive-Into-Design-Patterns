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
	if singletonInstance == nil {
		fmt.Println("Creating single instance now.")
		singletonInstance = &singleton{}
		fmt.Println(&singletonInstance)
		return singletonInstance
	}
	fmt.Println("Single instance already created.")
	return singletonInstance
}
