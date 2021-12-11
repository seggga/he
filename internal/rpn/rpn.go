package rpn

import "github.com/seggga/he/internal/domain"

func ConvertToRPN(query []domain.Token) []domain.Token {

	var rpn, stack []domain.Token
	var stackPriority, lexPriority int

	for _, lex := range query {
		// операнд сразу заносится в выходную строку
		if isOperand(lex) {
			rpn = append(rpn, lex)
			continue
		}

		// открывающая скобка сразу идет в стек
		if lex.Lexema == "(" {
			stack = append(stack, lex)
			continue
		}

		// закрывающая скобка извлекает все элементы из стека до символа "("
		if lex.Lexema == ")" {
			for i := len(stack) - 1; i >= 0; i -= 1 {
				stackLex := stack[i]
				stack = stack[:i]

				if stackLex.Lexema == "(" {
					break
				}

				rpn = append(rpn, stackLex)
			}
			continue
		}

		// оператор участвует в ветвлении
		if isOperator(lex) {
			lexPriority = getPriority(lex)

			// стек пуст - записываем знак операции в стек
			if len(stack) == 0 {
				stack = append(stack, lex)
				continue
			}

			// стек не пуст, извлекаем все операторы из стека, чей приоритет >= приоритету токена
			// затем помещаем токен в стек
			for i := len(stack) - 1; i >= 0; i -= 1 {

				stackLex := stack[i]
				stackPriority = getPriority(stackLex)

				// приоритет у токена больше, чем у вершины стека - кладем токен в стек
				if lexPriority > stackPriority {
					stack = append(stack, lex)
					break
				}

				stack = stack[:i]

				rpn = append(rpn, stackLex)
			}

			// если из стека извлекли все операторы, то токен заносим в стек
			if len(stack) == 0 {
				stack = append(stack, lex)
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

func isOperator(lex token.Lexema) bool {

	switch lex.Token {
	case "COMP":
		return true
	case "AND", "OR":
		return true
	default:
		return false

	}
}

func isOperand(lex token.Lexema) bool {

	if isOperator(lex) {
		return false
	}

	switch lex.Token {
	case "PAREN":
		return false
	default:
		return true
	}
}

func getPriority(lex token.Lexema) int {
	switch lex.Litera {
	case "(":
		return 0
	case ")":
		return 1
	case "AND", "OR":
		return 2
	case ">", "<", "==", ">=", "<=", "=":
		return 3
	default:
		return 100 // never gonna be returned
	}
}
