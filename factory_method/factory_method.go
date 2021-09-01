package main

import "fmt"

type Factory interface {
	FactoryMethod() Product
}
type Product interface {
	Use()
}
type ConcreteProductA struct {

}
type ConcreteProductB struct {

}

func (p *ConcreteProductA) Use() {
	fmt.Println("Product Use A...")
}
func (p *ConcreteProductB) Use() {
	fmt.Println("Product Use B...")
}
type ConcreateFactoryA struct {

}
type ConcreateFactoryB struct {
	
}

func (ConcreateFactoryB) FactoryMethod() Product {
	return &ConcreteProductB{}
}

func (f *ConcreateFactoryA) FactoryMethod() Product {
	return &ConcreteProductA{}
}




