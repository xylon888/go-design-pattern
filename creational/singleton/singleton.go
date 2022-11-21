package singleton

import (
	"fmt"
	"sync"
)

type singleton struct{}

var (
	instance *singleton
	once     sync.Once
)

func getInstance() *singleton {
	if instance == nil {
		once.Do(func() {
			fmt.Println("Creating single instance now")
			instance = &singleton{}
		})
	} else {
		fmt.Println("Single instance already created")
	}
	return instance
}
