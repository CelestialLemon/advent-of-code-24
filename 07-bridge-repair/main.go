package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)


type Equation struct {
	result int;
	operands []int;
}

type Stack []int;

func StackPush(s Stack, v int) Stack {
	return append(s, v);
}

func StackPop(s Stack) (Stack, int) {
	l := len(s);
	v := s[l - 1];
	return s[0 : l - 1], v;
}

type StringStack []string;

func StringStackPush(s StringStack, v string) StringStack {
	return append(s, v);
}

func StringStackPop(s StringStack) (StringStack, string) {
	l := len(s);
	v := s[l - 1];
	return s[0 : l - 1], v;
}

func readEquations(filePath string) ([]Equation, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var equations []Equation
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ": ")
		if len(parts) != 2 {
			return nil, fmt.Errorf("unexpected line format: %s", line)
		}
		result, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, err
		}

		operandStrings := strings.Split(parts[1], " ")
		var operands []int
		for _, operandString := range operandStrings {
			operand, err := strconv.Atoi(operandString)
			if err != nil {
				return nil, err
			}
			operands = append(operands, operand)
		}

		equations = append(equations, Equation{
			result: result,
			operands: operands,
		})
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return equations, nil
}

/**
	numOfOperators - length of the permutation
	operators = all the operator each place in a permutation can contain
 */
func getAllOperatorPermutations(numOfOperators int, operators []string ) [][]string {
	nOperators := len(operators);
	nPermutations := math.Pow(float64(nOperators), float64(numOfOperators));
	result := [][]string{};

	for i := 0; i < int(nPermutations); i++ {
		permutation := []string{}
		for f := i; f > 0; f /= nOperators {
			mod := f % nOperators;
			permutation = append(permutation, operators[mod]);
		}

		for len(permutation) != numOfOperators {
			permutation = append(permutation, operators[0]);
		}

		result = append(result, permutation);
	}

	return result;
}

func isEquationSatisfiable (eq Equation, operators []string) bool {
	allOperatorPerms := getAllOperatorPermutations(len(eq.operands) - 1, operators);
	for _, perm := range allOperatorPerms {
		
		// create operator stack
		operatorStack := StringStack{};

		for _, operator := range perm {
			operatorStack = StringStackPush(operatorStack, operator);
		}

		// create operand stack
		operandStack := Stack{};
		for i := len(eq.operands) - 1; i >= 0; i-- {
			operandStack = StackPush(operandStack, eq.operands[i]);
		}

		for len(operatorStack) > 0 {
			var a int;
			var b int;
			var operator string;

			operatorStack, operator = StringStackPop(operatorStack);
			operandStack, a = StackPop(operandStack);
			operandStack, b = StackPop(operandStack);


			result := 0;
			if operator == "+" {
				result = a + b;
			} else if operator == "*" {
				result = a * b;
			} else if operator == "||" {
				concatResult, err := strconv.Atoi(strconv.Itoa(a) + strconv.Itoa(b));
				if err != nil {
					fmt.Println(err);
					return false;
				}
				result = concatResult;
			}

			operandStack = StackPush(operandStack, result);
		}

		// only result will remain
		_, equationResult := StackPop(operandStack);

		if (equationResult == eq.result) {
			return true;
		}
	}

	return false;
}


func part1() {
	equations, err := readEquations("input.txt");

	if err != nil {
		fmt.Println(err);
	}

	result := 0
	for _, eq := range equations {
		if isEquationSatisfiable(eq, []string{"+", "*"}) {
			result += eq.result;
		} 
	}

	fmt.Println(result);
}

func part2() {
	equations, err := readEquations("input.txt");

	if err != nil {
		fmt.Println(err);
	}

	result := 0
	for _, eq := range equations {
		if isEquationSatisfiable(eq, []string{"+", "*", "||"}) {
			result += eq.result;
		} 
	}

	fmt.Println(result);
}

func main() {
	part2();
}
