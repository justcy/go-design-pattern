package main

import "fmt"

type Command interface {
	Execute()
}
type InvokerIface interface {
	Call()
}
type Invoker struct {
	com Command
}

func (i *Invoker) Call() {
	fmt.Println("invoker impl InvokerIface")
	i.com.Execute()
}

type ConcreateCommandA struct {
	rec *Receiver
}

func (c *ConcreateCommandA) Execute() {
	fmt.Println("Command A execute ")
	c.rec.OperationA()
}
func NewConcreateCommandA(receiver *Receiver) *ConcreateCommandA {
	return &ConcreateCommandA{
		rec: receiver,
	}
}

type ConcreateCommandB struct {
	rec *Receiver
}

func (c *ConcreateCommandB) Execute() {
	fmt.Println("Command B execute ")
	c.rec.OperationB()
}
func NewConcreateCommandB(receiver *Receiver) *ConcreateCommandB {
	return &ConcreateCommandB{
		rec: receiver,
	}
}

type Receiver struct {
}

func (r *Receiver) OperationA() {
	fmt.Println("Operation A ")
}
func (r *Receiver) OperationB() {
	fmt.Println("Operation B ")
}
func NewRecirver() *Receiver {
	return &Receiver{}
}
func NewInvoker(command Command) *Invoker {
	return &Invoker{
		com: command,
	}
}
