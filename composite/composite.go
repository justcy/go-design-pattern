package main

import "fmt"

type Component interface {
	Add(component Component)
	Remove(component Component)
	getChild(index int) Component
	Operation()
}

type Leaf struct {
	name string
}

func (l *Leaf) Add(component Component) {
	fmt.Printf("I'm a Leaf name:%s do nothing\n",l.name)
}

func (l *Leaf) Remove(component Component) {
	fmt.Printf("I'm a Leaf name:%s do nothing\n",l.name)
}

func (l *Leaf) getChild(index int) Component {
	fmt.Printf("I'm a Leaf name:%s do nothing\n",l.name)
	return nil
}

func (l  *Leaf) Operation() {
	fmt.Printf("I'm a Leaf name:%s do nothing\n",l.name)
}

func NewLeaf(name string) Component {
	return &Leaf{
		name:name,
	}
}
type Composite struct {
	components []Component
	name string
}

func (c *Composite) Add(component Component) {
	c.components  = append(c.components, component)
}

func (c *Composite) Remove(component Component) {
	for key, item := range c.components {
		if item == component{
			c.components = append(c.components[0:key],c.components[key+1:]...)
		}
	}
}

func (c *Composite) getChild(index int) Component {
	if index < 0 || index >len(c.components){
		return nil
	}
	return c.components[index]
}

func (c *Composite) Operation() {
	fmt.Printf("I'm a dir name:%s\n",c.name)
	for _, value := range c.components {
		value.Operation()
	}
}
func NewComposite(name string) Component {
	return &Composite{
		name:name,
	}
}
