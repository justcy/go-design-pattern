package main

import "fmt"

type Mediator interface {
	Operation(uid int, msg string)
	Register(uid int,colleage *Colleague)
}
type ConcreteMediator struct {
	name string
	usermap map[int]Colleague
}

func (c *ConcreteMediator) Operation(uid int, msg string) {
	colleague := c.usermap[uid]
	if colleague == nil{
		fmt.Println("can not find this colleague")
	}
	colleague.ReceiveMsg(msg)

}
func (c *ConcreteMediator) Register(uid int,colleague Colleague) {
	temp := c.usermap[uid]
	if temp == nil{
		c.usermap[uid] = colleague
	}
	colleague.SetMediator(c)

}
func NewConcreteMediator(name string) *ConcreteMediator{
	return &ConcreteMediator{
		name:name,
		usermap:make(map[int]Colleague),
	}
}

type Colleague interface {
	ReceiveMsg(str string)
	SendMsg(uid int,msg string)
	SetMediator(mediator *ConcreteMediator)
	GetName() string
}
type ConcreteColleagueA struct {
	mediator *ConcreteMediator
	name string
}

func (c *ConcreteColleagueA) GetName() string {
	return c.name
}

func (c *ConcreteColleagueA) ReceiveMsg(str string) {
	fmt.Printf("Receive message from Mediator %s , recive content:%s\n",c.mediator.name,str)
}

func (c *ConcreteColleagueA) SendMsg(uid int,msg string) {
	fmt.Printf("Send message from %s to %d,content:%s\n",c.name,uid,msg)
	c.mediator.Operation(uid,msg)
}

func (c *ConcreteColleagueA) SetMediator(m *ConcreteMediator) {
	c.mediator = m
}
func NewConcreteColleagueA(name string) Colleague{
	return &ConcreteColleagueA{
		name:name,
		mediator:nil,
	}
}
type ConcreteColleagueB struct {
	mediator *ConcreteMediator
	name string
}

func (c *ConcreteColleagueB) GetName() string {
	return c.name
}

func (c *ConcreteColleagueB) ReceiveMsg(str string) {
	fmt.Printf("Receive message from Mediator %s , recive content:%s\n",c.mediator.name,str)
}

func (c *ConcreteColleagueB) SendMsg(uid int,msg string) {
	fmt.Printf("Send message from %s to %d,content:%s\n",c.name,uid,msg)
	c.mediator.Operation(uid,msg)
}

func (c *ConcreteColleagueB) SetMediator(m *ConcreteMediator) {
	c.mediator = m
}
func NewConcreteColleagueB(name string) Colleague {
	return &ConcreteColleagueB{
		name:name,
		mediator:nil,
	}
}


