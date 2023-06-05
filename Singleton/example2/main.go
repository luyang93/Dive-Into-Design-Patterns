package main

import (
	"fmt"
	"sync"
)

type single struct{}

var singleInstance *single

var once sync.Once

func GetInstance() *single {
	once.Do(func() {
		singleInstance = &single{}
	})
	fmt.Println(&singleInstance)
	return singleInstance
}

func main() {
	for i := 0; i < 100; i++ {
		go GetInstance()
	}
}
