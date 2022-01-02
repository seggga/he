package rpn

import (
	"errors"
	"strings"

	"github.com/seggga/he/internal/domain"
)

type ConditionCheck struct {
	condition []domain.Token
}

// Init creates a new ConditionCheck structure
func (c *ConditionCheck) Init(whereTokens []domain.Token) {
	c.condition = convertToRPN(whereTokens)
}

func (c ConditionCheck) Check(head, row []string) (bool, error) {
	withValues := insertValues(head, row, c.condition)
	result, err := calculateRPN(withValues)
	if err != nil {
		return false, err
	}
	return result, nil
}

// convertToRPN changes tokens sequence to produce RPN
// has reference in NewConditionCheck function
func convertToRPN(condition []domain.Token) []domain.Token {
	var rpn, stack []domain.Token
	for _, token := range condition {
		// операнд сразу заносится в выходную строку
		if isOperand(token) {
			rpn = append(rpn, token)
			continue
		}

		// открывающая скобка сразу идет в стек
		if string(token.Lexema) == "(" {
			stack = append(stack, token)
			continue
		}

		// закрывающая скобка извлекает все элементы из стека до символа "("
		if string(token.Lexema) == ")" {
			for i := len(stack) - 1; i >= 0; i -= 1 {
				stackLex := stack[i]
				stack = stack[:i]

				if string(stackLex.Lexema) == "(" {
					break
				}

				rpn = append(rpn, stackLex)
			}
			continue
		}

		// оператор участвует в ветвлении
		if isOperator(token) {
			// если стек пуст - записываем знак операции в стек
			if len(stack) == 0 {
				stack = append(stack, token)
				continue
			}

			// стек не пуст, извлекаем все операторы из стека, чей приоритет >= приоритету токена
			// затем помещаем токен в стек
			for i := len(stack) - 1; i >= 0; i -= 1 {
				stackToken := stack[i]
				// приоритет у токена больше, чем у вершины стека - кладем токен в стек
				if token.Priority > stackToken.Priority {
					stack = append(stack, token)
					break
				}
				stack = stack[:i]
				rpn = append(rpn, stackToken)
			}

			// если из стека извлекли все операторы, то токен заносим в стек
			if len(stack) == 0 {
				stack = append(stack, token)
			}

		} // if isOperator(token)
	} // for _, token := range query

	// все токены просмотрены, надо опустошить стек
	for i := len(stack) - 1; i >= 0; i -= 1 {
		stackLex := stack[i]
		stack = stack[:i]
		rpn = append(rpn, stackLex)
	}
	return rpn
}

func isOperator(token domain.Token) bool {
	if token.TokenType == "operator" {
		return true
	}
	return false
}

func isOperand(token domain.Token) bool {
	if token.TokenType == "column name" || token.TokenType == "string" || token.TokenType == "integral" {
		return true
	}
	return false
}

// insertValues replases column names with their values
func insertValues(head, row []string, condition []domain.Token) []domain.Token {

	// create temp map to simplify operation
	valuesMap := make(map[string]string)
	for i := range head {
		valuesMap[head[i]] = row[i]
	}
	// create a temp slice to leave the original condition slice unchanged
	tempSlice := make([]domain.Token, len(condition))
	copy(tempSlice, condition)

	// change Lexema on elements with token type "column name"
	for i, v := range tempSlice {
		if v.TokenType == "column name" {
			tempSlice[i].Lexema = []byte(valuesMap[string(v.Lexema)])
		}
	}
	return tempSlice
}

var (
	error1 error = errors.New("error in expression, probably wrong query")
	error2 error = errors.New("the result was not obtained, probably wrong query")
)

// CalculateRPN - implements the mathematical calculation of given slice
// represented as reverse polish notation (RPN)
func calculateRPN(rpn []domain.Token) (bool, error) {

	for i := 0; i < len(rpn); {

		// if !token.IsOperand(&rpn[i]) {
		if !isOperand(rpn[i]) {
			i += 1
			continue
		}
		arg1 := rpn[i]

		// if !token.IsOperand(&rpn[i+1]) {
		if !isOperand(rpn[i+1]) {
			i += 1
			continue
		}
		arg2 := rpn[i+1]

		// if !token.IsOperator(&rpn[i+2]) {
		if !isOperator(rpn[i+2]) {
			i += 1
			continue
		}
		op := rpn[i+2]

		result, err := solveExpression(arg1, arg2, op)
		if err != nil {
			return false, err
		}

		// remove 3 elements from the rpn and move 'result' on that place
		newRPN := rpn[:i]
		newRPN = append(newRPN, result)
		newRPN = append(newRPN, rpn[i+3:]...)
		rpn = newRPN

		// start next iteration from '0' element in updated 'rpn'
		i = 0

		if len(rpn) == 1 {
			return string(rpn[0].Lexema) == "TRUE", nil
		}
	}

	return false, error2
}

// solveExpression calculates a given expression.
// As a result it makes token.Lexemma{Token: "IDENT", Litera: "TRUE"/"FALSE"}
// The main goal is to pass 'Litera' field with "TRUE" or "FALSE" value.
func solveExpression(arg1, arg2, op domain.Token) (domain.Token, error) {

	var result bool

	switch strings.ToLower(string(op.Lexema)) {
	case "=":
		result = string(arg1.Lexema) == string(arg2.Lexema)
	case ">=":
		result = string(arg1.Lexema) >= string(arg2.Lexema)
	case "<=":
		result = string(arg1.Lexema) <= string(arg2.Lexema)
	case ">":
		result = string(arg1.Lexema) > string(arg2.Lexema)
	case "<":
		result = string(arg1.Lexema) < string(arg2.Lexema)
	case "and":

		var operand1, operand2 bool

		// check if arg1 is 'true' / 'false'
		if string(arg1.Lexema) == "TRUE" {
			operand1 = true
		} else if string(arg1.Lexema) == "FALSE" {
			operand1 = false
		} else {
			return domain.Token{}, error1
		}

		// check if arg2 is 'true' / 'false'
		if string(arg2.Lexema) == "TRUE" {
			operand2 = true
		} else if string(arg2.Lexema) == "FALSE" {
			operand2 = false
		} else {
			return domain.Token{}, error1
		}

		result = operand1 && operand2

	case "or":

		var operand1, operand2 bool

		// check if arg1 is 'true' / 'false'
		if string(arg1.Lexema) == "TRUE" {
			operand1 = true
		} else if string(arg1.Lexema) == "FALSE" {
			operand1 = false
		} else {
			return domain.Token{}, error1
		}

		// check if arg2 is 'true' / 'false'
		if string(arg2.Lexema) == "TRUE" {
			operand2 = true
		} else if string(arg2.Lexema) == "FALSE" {
			operand2 = false
		} else {
			return domain.Token{}, error1
		}

		result = operand1 || operand2
	}

	if result {
		return domain.Token{Lexema: []byte("TRUE"), TokenType: "string"}, nil
	}
	return domain.Token{Lexema: []byte("FALSE"), TokenType: "string"}, nil
}
