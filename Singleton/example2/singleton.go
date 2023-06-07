package main

import (
	"fmt"
	"sync"
)

var (
	once              sync.Once
	singletonInstance *singleton
)

type singleton struct{}

func GetInstance() *singleton {
	once.Do(func() {
		singletonInstance = &singleton{}
	})
	fmt.Println(singletonInstance)
	return singletonInstance
}
