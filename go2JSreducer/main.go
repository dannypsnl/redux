package main

import (
	"fmt"
	"os"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/dannypsnl/redux/go2JSreducer/parser"
)

type TreeShapeListener struct {
	*parser.BaseReduxListener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (this *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println(ctx.GetText())
}

func main() {
	fmt.Println("translating")
	input, _ := antlr.NewFileStream(os.Args[1])
	lexer := parser.NewReduxLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewReduxParser(stream)
	p.BuildParseTrees = true
	tree := p.Prog()
	antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)
}
