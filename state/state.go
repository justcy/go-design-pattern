package main

import "fmt"

type State interface {
	Handle(context *Context)
}
type ConcreteStateA struct {

}

func (c *ConcreteStateA) Handle(context *Context) {
	fmt.Println("doing something in state A \n done. \n change state to B ")
	context.ChangeState(NewStateB())
}
func NewStateA() State {
	return &ConcreteStateA{}
}
func NewStateB() State {
	return &ConcreteStateB{}
}

type ConcreteStateB struct {

}

func (c *ConcreteStateB) Handle(context *Context) {
	fmt.Println("doing something in state B \n done. \n change state to A ")
	context.ChangeState(NewStateA())
}
type ContextInface interface {
	ChangeState(state int)
	Request()
}
type Context struct {
	state State
}

func (c *Context) ChangeState(state State) {
	c.state = state
}

func (c *Context) Request() {
	c.state.Handle(c)
}

func NewContext() *Context{
	return &Context{
		state:NewStateA(),
	}
}