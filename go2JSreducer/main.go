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

func (s *TreeShapeListener) EnterReducer(ctx *parser.ReducerContext) {
	fmt.Println("reducer found!", ctx.ID(0).GetText())
	fmt.Println("function", ctx.ID(0).GetText(), "(state =", ctx.GetInitialState().GetText(), ",action) {")
	fmt.Printf("  switch (%s.type) {\n", ctx.ID(4).GetText())
}
func (s *TreeShapeListener) EnterCases(ctx *parser.CasesContext) {
	fmt.Println("  case", ctx.Expr(0).GetText(), ":")
	fmt.Println("    return", ctx.Expr(1).GetText())
}
func (s *TreeShapeListener) ExitReducer(ctx *parser.ReducerContext) {
	fmt.Println("  }\n}")
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
