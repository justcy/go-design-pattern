package main

import (
	"strconv"
	"strings"
)
type Expression interface {
	Interpreter() int
}

type TerminalExpression struct {
	val int
}

func (t *TerminalExpression) Interpreter() int {
	return t.val
}

type AddExpression struct {
	left, right Expression
}

type SubExpression struct {
	left, right Expression
}

type Parser struct {
	exp   []string
	index int
	prev  Expression
}

func (n *AddExpression) Interpreter() int {
	return n.left.Interpreter() + n.right.Interpreter()
}
func (n *SubExpression) Interpreter() int {
	return n.left.Interpreter() - n.right.Interpreter()
}

func (p *Parser) Parse(exp string) {
	p.exp = strings.Split(exp, " ")

	for {
		if p.index >= len(p.exp) {
			return
		}
		switch p.exp[p.index] {
		case "+":
			p.prev = p.newAddExpression()
		case "-":
			p.prev = p.newSubExpression()
		default:
			p.prev = p.newTerminalExpression()
		}
	}
}
func (p *Parser) newAddExpression() *AddExpression {
	p.index++
	return &AddExpression{
		left:  p.prev,
		right: p.newTerminalExpression(),
	}
}

func (p *Parser) newSubExpression() *SubExpression {
	p.index++
	return &SubExpression{
		left:  p.prev,
		right: p.newTerminalExpression(),
	}
}

func (p *Parser) newTerminalExpression() *TerminalExpression {
	v, _ := strconv.Atoi(p.exp[p.index])
	p.index++
	return &TerminalExpression{
		val: v,
	}
}
func (p *Parser) Result() Expression {
	return p.prev
}
