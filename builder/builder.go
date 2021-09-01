package main

import "fmt"

type Builder interface {
	BuilderPartA()
	BuilderPartB()
	BuilderPartC()
	GetResult()
}
type ConcreteBuilder struct {
	result string
}

func (c *ConcreteBuilder) BuilderPartA() {
	fmt.Println("Build Part A")
	c.result += "A"
}

func (c *ConcreteBuilder) BuilderPartB() {
	fmt.Println("Build Part B")
	c.result += "B"
}

func (c *ConcreteBuilder) BuilderPartC() {
	fmt.Println("Build Part C")
	c.result += "C"
}

func (c *ConcreteBuilder) GetResult() {
	fmt.Printf("Get Result %s",c.result)
}

type Director struct {
	builder Builder
}

func NewDirector(builder Builder) *Director {
	return &Director{
		builder:builder,
	}
}
func (d *Director) Construct() {
	d.builder.BuilderPartA()
	d.builder.BuilderPartB()
	d.builder.BuilderPartC()
}


