package main

import "fmt"

//Target 适配的目标接口
type Target interface {
	Request() string
}
//Adaptee 是被适配的目标接口
type Adapter interface {
	SpecificRequest() string
}

type Adaptee struct {
}

// adapter 是转换Adaptee为Target接口的适配器
type adapter struct {
	Adapter
}
//Request 实现Target接口
func (a *adapter) Request() string {
	return a.SpecificRequest()
}
//SpecificRequest 是目标类的一个方法
func (a *Adaptee) SpecificRequest() string {
	fmt.Println("Specific Request method called")
	return "adaptee method"
}
//NewAdaptee 是被适配接口的工厂函数
func NewAdaptee() Adapter {
	return &Adaptee{}
}
//NewAdapter 是Adapter的工厂函数
func NewAdapter(adaptee Adapter) Target {
	return &adapter{
		Adapter: adaptee,
	}
}
