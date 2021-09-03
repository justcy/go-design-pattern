package main

import "fmt"

func main() {
	manager := NewPrototypeManager()
	t1 := &ConcretePrototype{
		name:"type1",
	}
	manager.Set("t1",t1)

	t2 := manager.Get("t1")
	t3 := t2.Clone()
	if t2 == t3 {
		fmt.Println("error! get clone not working")
	} else {
		fmt.Println("get clone working")
	}

	c := manager.Get("t1").Clone()

	t4 := c.(*ConcretePrototype)
	if t4.name != "type1" {
		fmt.Println("error")
	}
}
