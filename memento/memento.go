package main

import "fmt"

type Originator struct {
	state int
}

func NewOriginator() *Originator  {
	return &Originator{
		state:1,
	}
}

func (o *Originator) CreateMemento() Memento {
	return &GameMemento{
		state:o.state,
	}
}
func (o *Originator) RestoreMemento(m Memento) {
	o.state = m.GetState()
}
func (o *Originator) Change(state int) {
	o.state = state
}
type Memento interface {
	SetState(state int)
	GetState() int
}
type GameMemento struct {
	state int
}

func (g *GameMemento) SetState( state int) {
	g.state =state
}

func (g *GameMemento) GetState() int {
	return g.state
}
func (o *Originator) Status()  {
	fmt.Printf("Current State:%d\n", o.state)
}

