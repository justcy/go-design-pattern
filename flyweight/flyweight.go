package main

import "fmt"

type Flyweight interface {
	Operation()
}
type ConcreteFlyweight struct {
	name string
}

func (c *ConcreteFlyweight) Operation() {
	fmt.Printf("flyweight[%s] impl concrete \n",c.name)
}

type UnsharedConcreteFlyweight struct {

}

func (u *UnsharedConcreteFlyweight) Operation() {
	fmt.Println("flyweight impl UnsharedConcrete ")
}

type FlyweightFactory struct {
	maps map[string] *ConcreteFlyweight
}

var factory *FlyweightFactory

func NewFlyweightFactory()  *FlyweightFactory {
	if factory == nil{
		factory = &FlyweightFactory{
			maps: make(map[string]*ConcreteFlyweight ),
		}
	}
	return factory
}

func NewConcreteFlyweight(name string ) *ConcreteFlyweight  {
	return &ConcreteFlyweight{
		name:name,
	}
}
func NewUnshargConcreteFlyweight() *UnsharedConcreteFlyweight  {
	return &UnsharedConcreteFlyweight{}
}
func (f *FlyweightFactory) Get(name string) *ConcreteFlyweight {
	flyweight := f.maps[name]
	if flyweight == nil{
		flyweight = NewConcreteFlyweight(name)
		f.maps[name] = flyweight
	}else{
		fmt.Println("already exist "+name+",use it !")
	}
	return flyweight
}