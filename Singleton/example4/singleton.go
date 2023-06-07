package main

import (
	"fmt"
	"sync"
)

var (
	wg                sync.WaitGroup
	singletonInstance *singleton
)

type singleton struct{}

func GetInstance() *singleton {
	defer wg.Done()
	if singletonInstance == nil {
		fmt.Println("Creating single instance now.")
		singletonInstance = &singleton{}
		return singletonInstance
	}
	fmt.Println("Single instance already created.")
	return singletonInstance
}
