package main

import "fmt"

type Abstraction interface {
	Operation()

}
type RefinedAbstraction interface {
	Abstraction
}
type Implementor interface {
	OperationImpl()
}

type ConcreteImplementorA struct {

}

func NewImplementorA() Implementor  {
	return &ConcreteImplementorA{}
}
func (c *ConcreteImplementorA) OperationImpl() {
	fmt.Println("This is Implementor A")
}

type ConcreteImplementorB struct {

}

func NewImplementorB() Implementor  {
	return &ConcreteImplementorB{}
}

func (c *ConcreteImplementorB) OperationImpl() {
	fmt.Println("This is Implementor B")
}

type CommonImplementor struct {
	method Implementor
}

func (c *CommonImplementor) Operation() {
	fmt.Println("CommonImplementor1 Call")
	c.method.OperationImpl()
}

func NewCommonImplementor(method Implementor)  *CommonImplementor {
	return &CommonImplementor{
		method:method,
	}
}

type CommonImplementor2 struct {
	method Implementor
}

func (c *CommonImplementor2) Operation() {
	fmt.Println("CommonImplementor2 Call")
	c.method.OperationImpl()
}

func NewCommonImplementor2(method Implementor)  *CommonImplementor2 {
	return &CommonImplementor2{
		method:method,
	}
}
