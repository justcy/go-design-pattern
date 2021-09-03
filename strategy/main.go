package main

func main() {
	context := NewContext()
	strategyA := NewConcreteStrategyA()
	context.SetStrategy(strategyA)
	context.strategy.Algorithm()

	strategyB := NewConcreteStrategyB()
	context.SetStrategy(strategyB)
	context.strategy.Algorithm()
}
