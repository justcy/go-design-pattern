package main

type Abstraction interface {
	Operation()

}
type RefinedAbstraction interface {
	Abstraction
}
type Implementor interface {
	RefinedAbstraction
}

type ConcreteImplementorA struct {

}
type ConcreteImplementorB struct {

}