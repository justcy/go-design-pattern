package main

func main() {
	var factory Factory
	product := factory.CreateProduct("A")
	product.Use()

	productB := factory.CreateProduct("B")
	productB.Use()
}
