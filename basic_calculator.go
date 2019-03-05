// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ Copyright (c) Michael Leahcim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2019-03-05 13:21 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@

package main

// Solving: https://leetcode.com/problems/basic-calculator/

// "1+1"
// "2-1+2"
// "(1+(4+5+2)-3)+(6+8)"
// ((((1+2))))

const OP_PLUS = -1
const OP_MINUS = -2
const OP_LEFT = -3
const OP_RIGHT = -4
const OP_NO = -5

func calculateTokenize(s string) []int {
	var prevNum int
	sTokenized := []int{}
	upNum := func(num int) {
		// is number?
		if len(sTokenized) > 0 && sTokenized[len(sTokenized)-1] >= 0 {
			// pop op
			prevNum, sTokenized = sTokenized[len(sTokenized)-1], sTokenized[:len(sTokenized)-1]
			sTokenized = append(sTokenized, prevNum*10+num)
			return
		}
		sTokenized = append(sTokenized, num)
	}
	for _, v := range s {
		switch v {
		case ' ':
			continue
		case '(':
			sTokenized = append(sTokenized, OP_LEFT)
		case ')':
			sTokenized = append(sTokenized, OP_RIGHT)
		case '+':
			sTokenized = append(sTokenized, OP_PLUS)
		case '-':
			sTokenized = append(sTokenized, OP_MINUS)
		case '0':
			upNum(0)
		case '1':
			upNum(1)
		case '2':
			upNum(2)
		case '3':
			upNum(3)
		case '4':
			upNum(4)
		case '5':
			upNum(5)
		case '6':
			upNum(6)
		case '7':
			upNum(7)
		case '8':
			upNum(8)
		case '9':
			upNum(9)
		default:
			panic("Not expected")
		}
	}
	return sTokenized
}

func calculate(s string) int {
	var op int
	var num int

	sTokenized := calculateTokenize(s)

	stack := []int{}
	op_stack := []int{OP_NO}
	for _, token := range sTokenized {
		switch token {
		case OP_LEFT:
			op_stack = append(op_stack, OP_NO)
		case OP_PLUS:
			op_stack = append(op_stack, OP_PLUS)
		case OP_MINUS:
			op_stack = append(op_stack, OP_MINUS)
		default:
			if token == OP_RIGHT {
				token, stack = stack[len(stack)-1], stack[:len(stack)-1]
			}
			op, op_stack = op_stack[len(op_stack)-1], op_stack[:len(op_stack)-1]
			switch op {
			case OP_NO:
				stack = append(stack, token)
			case OP_PLUS:
				num, stack = stack[len(stack)-1], stack[:len(stack)-1]
				stack = append(stack, token+num)
			case OP_MINUS:
				num, stack = stack[len(stack)-1], stack[:len(stack)-1]
				stack = append(stack, num-token)
			}
		}
	}
	return stack[0]

}
