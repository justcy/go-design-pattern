package main

func main() {
	builder := &ConcreteBuilder{}
	director := NewDirector(builder)
	director.Construct()
	builder.GetResult()
}
