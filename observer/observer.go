package main

import "fmt"

type Observer interface {
	Update(subject Subject)
}
type ConcreteObserver struct {
	name string
}

func (c *ConcreteObserver) Update(subject Subject) {
	fmt.Printf("Observer %s,Observed state change %d \n",c.name,subject.GetState())
}

type Subject interface {
	Attach(observer Observer)
	Detach(observer Observer)
	Notify()

	GetState() int
	SetState(state int)
}

type ConcreteSubject struct {
	objList []Observer
	state   int
}

func NewSubject() *ConcreteSubject {
	return &ConcreteSubject{
		objList: make([]Observer, 0),
		state:   0,
	}
}

func NewConcreteObserver(name string) Observer{
	return &ConcreteObserver{
		name:name,
	}
}
func (c *ConcreteSubject) Attach(observer Observer) {
	c.objList = append(c.objList, observer)
}

func (c *ConcreteSubject) Detach(observer Observer) {
	for key, obj := range c.objList {
		if obj == observer {
			c.objList = append(c.objList[:key], c.objList[key+1:]...)
		}
	}
}

func (c *ConcreteSubject) Notify() {
	for _, obj := range c.objList {
		obj.Update(c)
	}
}

func (c *ConcreteSubject) GetState() int {
	return c.state
}

func (c *ConcreteSubject) SetState(state int) {
	c.state = state
}
