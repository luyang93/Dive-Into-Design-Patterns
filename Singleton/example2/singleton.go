package main

import (
	"fmt"
	"sync"
)

var (
	singletonInstance *singleton
	once              sync.Once
	wg                sync.WaitGroup
)

type singleton struct{}

func GetInstance() *singleton {
	defer wg.Done()
	once.Do(func() {
		singletonInstance = &singleton{}
	})
	fmt.Println(singletonInstance)
	return singletonInstance
}
