package main

import "fmt"

//抽象工厂接口
type AbstractFactory interface {
	CreateProductA() AbstractProductA
	CreateProductB() AbstractProductB
}

//抽象产品A接口
type AbstractProductA interface {
	Use()
}
//抽象产品B接口
type AbstractProductB interface {
	Eat()
}

//抽象工厂实现1
type ConcreteFactory1 struct {
}

func (fc *ConcreteFactory1) CreateProductA() AbstractProductA {
	return &ProductA1{}
}

func (fc *ConcreteFactory1) CreateProductB() AbstractProductB {
	return &ProductB1{}
}
//抽象工厂实现2
type concreteFactory2 struct {
}

func (concreteFactory2) CreateProductA() AbstractProductA {
	return &ProductA2{}
}

func (concreteFactory2) CreateProductB() AbstractProductB {
	return &ProductB2{}
}

type ProductA1 struct {
}
type ProductA2 struct {
}

func (A2 *ProductA2) Use() {
	fmt.Println("Product A2 How to Use...")
}

func (A *ProductA1) Use() {
	fmt.Println("Product A1 How to Use...")
}

type ProductB1 struct {
}
type ProductB2 struct {
}

func (B *ProductB2) Eat() {
	fmt.Println("Product B2 Begin Eat...")
}

func (B *ProductB1) Eat() {
	fmt.Println("Product B1 Begin Eat...")
}
