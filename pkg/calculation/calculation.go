package calculation

import (
    "errors"
    "strconv"
    "strings"
    "unicode"
)

func Calc(expression string) (float64, error) {
    expression = strings.ReplaceAll(expression, " ", "")

    for _, r := range expression {
        if !unicode.IsDigit(r) && !strings.ContainsRune("+-*/().", r) {
            return 0, ErrInvalidExpression
        }
    }

    tokens, err := tokenize(expression)
    if err != nil {
        return 0, err
    }

    return evaluate(tokens)
}

func tokenize(expression string) ([]string, error) {
    var tokens []string
    var number strings.Builder

    for i, r := range expression {
        if unicode.IsDigit(r) || r == '.' {
            number.WriteRune(r)
        } else {
            if number.Len() > 0 {
                tokens = append(tokens, number.String())
                number.Reset()
            }

            if strings.ContainsRune("+-*/()", r) {
                tokens = append(tokens, string(r))
            } else {
                return nil, errors.New("unexpected character at position " + strconv.Itoa(i))
            }
        }
    }

    if number.Len() > 0 {
        tokens = append(tokens, number.String())
    }

    return tokens, nil
}

func evaluate(tokens []string) (float64, error) {
    var output []string
    var operators []string

    precedence := map[string]int{
        "+": 1, "-": 1,
        "*": 2, "/": 2,
    }

    for _, token := range tokens {
        if isNumber(token) {
            output = append(output, token)
        } else if token == "(" {
            operators = append(operators, token)
        } else if token == ")" {
            for len(operators) > 0 && operators[len(operators)-1] != "(" {
                output = append(output, operators[len(operators)-1])
                operators = operators[:len(operators)-1]
            }
            if len(operators) == 0 {
                return 0, errors.New("mismatched parentheses")
            }
            operators = operators[:len(operators)-1]
        } else {
            for len(operators) > 0 && precedence[operators[len(operators)-1]] >= precedence[token] {
                output = append(output, operators[len(operators)-1])
                operators = operators[:len(operators)-1]
            }
            operators = append(operators, token)
        }
    }

    for len(operators) > 0 {
        output = append(output, operators[len(operators)-1])
        operators = operators[:len(operators)-1]
    }

    return evalRPN(output)
}

func isNumber(token string) bool {
    _, err := strconv.ParseFloat(token, 64)
    return err == nil
}

func evalRPN(tokens []string) (float64, error) {
    var stack []float64

    for _, token := range tokens {
        if isNumber(token) {
            num, _ := strconv.ParseFloat(token, 64)
            stack = append(stack, num)
        } else {
            if len(stack) < 2 {
                return 0, errors.New("invalid expression")
            }
            b := stack[len(stack)-1]
            a := stack[len(stack)-2]
            stack = stack[:len(stack)-2]

            var result float64
            switch token {
            case "+":
                result = a + b
            case "-":
                result = a - b
            case "*":
                result = a * b
            case "/":
                if b == 0 {
                    return 0, errors.New("division by zero")
                }
                result = a / b
            }

            stack = append(stack, result)
        }
    }

    if len(stack) != 1 {
        return 0, errors.New("invalid expression")
    }

    return stack[0], nil
}
