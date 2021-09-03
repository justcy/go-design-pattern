package main

import "fmt"

type Conmponent interface {
	Operation()
}
type ConcreteComponent struct {
}

func (*ConcreteComponent) Operation() {
	fmt.Println("Operation Called ")
}
func NewConcreteComponent() Conmponent {
	return &ConcreteComponent{}
}

//type Decorator struct {
//	com Conmponent
//}
type ConcreteDecoratorA struct {
	com Conmponent
}

func NewConcreteDecoratorA(c Conmponent) Conmponent {
	return &ConcreteDecoratorA{
		com: c,
	}
}
func (c *ConcreteDecoratorA) AddBehavior() {
	fmt.Println("add Behavior AAAA")
}

func (c *ConcreteDecoratorA) Operation() {
	fmt.Println("ConcreteDecoratorA Operation Called ")
	c.AddBehavior()
}

type ConcreteDecoratorB struct {
	com Conmponent
}

func NewConcreteDecoratorB(c Conmponent) Conmponent {
	return &ConcreteDecoratorB{
		com: c,
	}
}
func (c *ConcreteDecoratorB) Operation() {
	fmt.Println("ConcreteDecoratorB Operation Called ")
	c.AddBehavior()
}

func (c *ConcreteDecoratorB) AddBehavior() {
	fmt.Println("add Behavior BBBB")
}
