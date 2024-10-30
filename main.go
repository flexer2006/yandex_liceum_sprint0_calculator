package main

import (
	"errors"
	"strconv"
	"unicode"
)

func Calc(expression string) (float64, error) {
	var nums []float64
	var ops []rune

	precedence := map[rune]int{
		'+': 1,
		'-': 1,
		'*': 2,
		'/': 2,
	}

	applyOp := func() error {
		if len(nums) < 2 || len(ops) == 0 {
			return errors.New("invalid expression")
		}
		b, a := nums[len(nums)-1], nums[len(nums)-2]
		nums = nums[:len(nums)-2]
		op := ops[len(ops)-1]
		ops = ops[:len(ops)-1]

		var res float64
		switch op {
		case '+':
			res = a + b
		case '-':
			res = a - b
		case '*':
			res = a * b
		case '/':
			if b == 0 {
				return errors.New("division by zero")
			}
			res = a / b
		default:
			return errors.New("invalid operator")
		}
		nums = append(nums, res)
		return nil
	}

	for i := 0; i < len(expression); i++ {
		ch := rune(expression[i])

		if unicode.IsDigit(ch) || ch == '.' {
			j := i
			for j < len(expression) && (unicode.IsDigit(rune(expression[j])) || rune(expression[j]) == '.') {
				j++
			}
			num, err := strconv.ParseFloat(expression[i:j], 64)
			if err != nil {
				return 0, err
			}
			nums = append(nums, num)
			i = j - 1
		} else if ch == '(' {
			ops = append(ops, ch)
		} else if ch == ')' {
			for len(ops) > 0 && ops[len(ops)-1] != '(' {
				if err := applyOp(); err != nil {
					return 0, err
				}
			}
			if len(ops) == 0 {
				return 0, errors.New("mismatched parentheses")
			}
			ops = ops[:len(ops)-1]
		} else if _, ok := precedence[ch]; ok {
			for len(ops) > 0 && precedence[ops[len(ops)-1]] >= precedence[ch] {
				if err := applyOp(); err != nil {
					return 0, err
				}
			}
			ops = append(ops, ch)
		} else if !unicode.IsSpace(ch) {
			return 0, errors.New("invalid character")
		}
	}

	for len(ops) > 0 {
		if err := applyOp(); err != nil {
			return 0, err
		}
	}

	if len(nums) != 1 {
		return 0, errors.New("invalid expression")
	}

	return nums[0], nil
}
