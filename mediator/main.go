package main

func main() {
	colleagueA := NewConcreteColleagueA("zhangsan")
	colleagueB := NewConcreteColleagueB("lisi")

	mediator := NewConcreteMediator("wangwu")

	mediator.Register(1,colleagueA)
	mediator.Register(2,colleagueB)

	colleagueA.SendMsg(2,"hello I'm "+colleagueA.GetName())
	colleagueB.SendMsg(1,"hello I'm "+colleagueB.GetName()+" Nice to meet you ")
	colleagueB.SendMsg(3,"hello I'm "+colleagueB.GetName()+" Nice to meet you ")
}
