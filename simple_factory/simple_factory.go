package main

import "fmt"

type Product interface {
	Use()
}
type ConcreteProductA struct {

}

func (p *ConcreteProductA) Use() {
	fmt.Println("Product A use")
}

type ConcreteProductB struct {

}

func (p *ConcreteProductB) Use() {
	fmt.Println("Product B use")
}

type Factory struct {

}

func (f Factory) CreateProduct(str string) Product  {
	if str == "A"{
		return &ConcreteProductA{}
	}else if str == "B"{
		return &ConcreteProductB{}
	}
	return &ConcreteProductA{}
}
