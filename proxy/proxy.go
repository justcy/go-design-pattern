package main

import "fmt"

type Subject interface {
	Request()
}

type RealSubject struct {

}

func (r *RealSubject) Request() {
	fmt.Println("real subject request do something..")
}

type ProxyIface interface {
	preRequest()
	afterRequest()
}
type Proxy struct {
	real RealSubject
}

func (p *Proxy) preRequest()  {
	fmt.Println("proxy pre request do something..")
}
func (p *Proxy) Request()  {
	p.preRequest()
	fmt.Println("proxy request do something..")
	p.real.Request()
	p.afterRequest()
}
func (p *Proxy) afterRequest()  {
	fmt.Println("proxy after request do something..")
}
