package main

import (
	"fmt"
	"sync"
)

type Singleton struct {
}

var (
	instance *Singleton
	once sync.Once
)

func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{}
	})
	return instance
}
func (s *Singleton) say()  {
	fmt.Println("say Hello ")
}