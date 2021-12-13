package rpn

import "github.com/seggga/he/internal/domain"

func ConvertToRPN(conditionTokens []domain.Token) []domain.Token {

	var rpn, stack []domain.Token
	// var stackPriority, tokenPriority int

	for _, token := range conditionTokens {
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
			// tokenPriority = getPriority(lex)

			// стек пуст - записываем знак операции в стек
			if len(stack) == 0 {
				stack = append(stack, token)
				continue
			}

			// стек не пуст, извлекаем все операторы из стека, чей приоритет >= приоритету токена
			// затем помещаем токен в стек
			for i := len(stack) - 1; i >= 0; i -= 1 {

				stackToken := stack[i]
				// stackPriority = stack[i].Priority

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
