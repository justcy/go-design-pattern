package main

func main() {
	var factory Factory
	factory = &ConcreateFactoryA{}
	product := factory.FactoryMethod()
	product.Use()

	factory = &ConcreateFactoryB{}
	product = factory.FactoryMethod()
	product.Use()
}

