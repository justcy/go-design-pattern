package main

func main() {
	var c Conmponent = &ConcreteComponent{}
	c.Operation()

	c = NewConcreteDecoratorA(c)
	c.Operation()

	c = NewConcreteDecoratorB(c)
	c.Operation()
}
