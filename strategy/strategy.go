package main

import "fmt"

type Strategy interface {
	Algorithm()
}

type ConcreteStrategyA struct {
}

func (a *ConcreteStrategyA) Algorithm() {
	fmt.Println("Use StrategyA First in First out!")
}
func NewConcreteStrategyA() *ConcreteStrategyA {
	return &ConcreteStrategyA{}
}

type ConcreteStrategyB struct {
}

func (b ConcreteStrategyB) Algorithm() {
	fmt.Println("Use StrategyB Last in First out !")
}

func NewConcreteStrategyB() *ConcreteStrategyB {
	return &ConcreteStrategyB{}
}

type Context struct {
	strategy Strategy
}

func (c *Context) SetStrategy(a Strategy) {
	c.strategy = a
}

func NewContext() *Context {
	return &Context{}
}
