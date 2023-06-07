package main

import (
	"sync"
)

type Singleton struct {
	ID int
	// 定义单例类的其他属性
}

var (
	instances []*Singleton
	mutex     sync.Mutex
)

// GetInstance 返回单例实例
func GetInstance(limit int) *Singleton {
	if limit <= 0 {
		// 不限制单例实例的数量
		return &Singleton{}
	}

	mutex.Lock()
	defer mutex.Unlock()

	if len(instances) < limit {
		// 创建新的单例实例
		instance := &Singleton{}
		instances = append(instances, instance)
		return instance
	}

	// Move the first instance to the end of the instances slice
	firstInstance := instances[0]
	instances = instances[1:]
	instances = append(instances, firstInstance)

	return firstInstance
}
