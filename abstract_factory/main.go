package main

func showUseAndEat(factory AbstractFactory)  {
	factory.CreateProductA().Use()
	factory.CreateProductB().Eat()
}
func main() {
	var factory AbstractFactory
	factory = &ConcreteFactory1{}
	showUseAndEat(factory)

	factory = &concreteFactory2{}
	showUseAndEat(factory)
}
