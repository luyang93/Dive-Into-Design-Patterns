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

// wrong example
func GetInstance() *singleton {
	defer wg.Done()
	singletonInstance = &singleton{}
	fmt.Println(&singletonInstance)
	return singletonInstance
}

// wrong example
func GetInstance2() *singleton {
	defer wg.Done()
	singletonInstance2 := &singleton{}
	fmt.Println(&singletonInstance2)
	return singletonInstance2
}
