package main

import "fmt"

type Handler interface {
	SetSuccessor(handler Handler)
	HandleRequest(req int)
}
type ConcreteHandler1 struct {
	successor Handler
}

func (c *ConcreteHandler1) SetSuccessor(handler Handler) {
	c.successor = handler
}

func (c *ConcreteHandler1) HandleRequest(req int) {
	if req >= 0 && req<10 {
		fmt.Printf("ConcreteHandler1 处理请求request = %d \n", req)
	}else if c.successor !=nil {
		c.successor.HandleRequest(req)
	}
}
func NewConcreteHandler1() *ConcreteHandler1{
	return &ConcreteHandler1{
		successor:nil,
	}
}
type ConcreteHandler2 struct {
	successor Handler
}

func NewConcreteHandler2() *ConcreteHandler2{
	return &ConcreteHandler2{
		successor:nil,
	}
}
func (c *ConcreteHandler2) SetSuccessor(handler Handler) {
	c.successor = handler
}

func (c *ConcreteHandler2) HandleRequest(req int) {
	if req >= 10 && req<20 {
		fmt.Printf("ConcreteHandler2 处理请求request = %d \n", req)
	}else if c.successor !=nil {
		c.successor.HandleRequest(req)
	}
}

type ConcreteHandler3 struct {
	successor Handler
}
func NewConcreteHandler3() *ConcreteHandler3{
	return &ConcreteHandler3{
		successor:nil,
	}
}

func (c *ConcreteHandler3) SetSuccessor(handler Handler) {
	c.successor = handler
}

func (c *ConcreteHandler3) HandleRequest(req int) {
	if req >= 20 && req<30 {
		fmt.Printf("ConcreteHandler3 处理请求request = %d \n", req)
	}else if c.successor !=nil {
		c.successor.HandleRequest(req)
	}
}


