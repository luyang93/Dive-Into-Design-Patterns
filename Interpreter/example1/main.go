package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Expression interface {
	Interpret()
}

type PrintExpression struct {
	Message string
}

func (pe *PrintExpression) Interpret() {
	fmt.Println(pe.Message)
}

type RepeatExpression struct {
	RepeatCount int
	Expression  Expression
}

func (re *RepeatExpression) Interpret() {
	for i := 0; i < re.RepeatCount; i++ {
		re.Expression.Interpret()
	}
}

type ReverseExpression struct {
	Message string
}

func (re *ReverseExpression) Interpret() {
	runes := []rune(re.Message)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	fmt.Println(string(runes))
}

func main() {
	command := "REPEAT 3 TIMES: REVERSE Hello"
	words := strings.Split(command, " ")

	if words[0] == "REPEAT" {
		repeatCount, err := strconv.Atoi(words[1])
		if err != nil {
			panic(err)
		}

		var expression Expression
		if words[3] == "PRINT" {
			expression = &PrintExpression{Message: words[4]}
		} else if words[3] == "REVERSE" {
			expression = &ReverseExpression{Message: words[4]}
		} else {
			panic("Unsupported command: " + words[3])
		}

		repeatExpression := &RepeatExpression{RepeatCount: repeatCount, Expression: expression}
		repeatExpression.Interpret()
	}
}
