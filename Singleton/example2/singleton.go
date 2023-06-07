package main

import (
	"fmt"
	"sync"
)

var (
	once              sync.Once
	wg                sync.WaitGroup
	singletonInstance *singleton
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
