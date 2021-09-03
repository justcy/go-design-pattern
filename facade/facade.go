package main

import "fmt"

type FacadeInterface interface {
	WrapOperation()
}

type Facade struct {
	A AbstractSystemA
	B AbstractSystemB
	C AbstractSystemC
}

func NewFacade() FacadeInterface {
	return &Facade{
		A: NewSystemA(),
		B: NewSystemB(),
		C: NewSystemC(),
	}
}
func (f *Facade) WrapOperation() {
	f.A.OperationA()
	f.B.OperationB()
	f.C.OperationC()
}

type AbstractSystemA interface {
	OperationA()
}
type SystemA struct {
}

func (a *SystemA) OperationA() {
	fmt.Println("System A operation ..")
}
func NewSystemA() AbstractSystemA {
	return &SystemA{}
}

type AbstractSystemB interface {
	OperationB()
}
type SystemB struct {
}

func NewSystemB() AbstractSystemB {
	return &SystemB{}
}
func (b *SystemB) OperationB() {
	fmt.Println("System B operation ..")
}

type AbstractSystemC interface {
	OperationC()
}
type SystemC struct {
}

func (c *SystemC) OperationC() {
	fmt.Println("System C operation ..")
}
func NewSystemC() AbstractSystemC {
	return &SystemC{}
}
