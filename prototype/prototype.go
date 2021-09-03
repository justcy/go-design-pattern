package main

//Cloneable 是原型对象需要实现的接口
type IPrototype interface {
	Clone() IPrototype
}

type PrototypeManager struct {
	prototypes map[string]IPrototype
}

func NewPrototypeManager() *PrototypeManager {
	return &PrototypeManager{
		prototypes: make(map[string]IPrototype),
	}
}

func (p *PrototypeManager) Get(name string) IPrototype {
	return p.prototypes[name].Clone()
}

func (p *PrototypeManager) Set(name string, prototype IPrototype) {
	p.prototypes[name] = prototype
}

type ConcretePrototype struct {
	name string
}

func (c *ConcretePrototype) Clone() IPrototype {
	tc := *c
	return &tc
}


