package jsonparser

import (
	"strconv"
	"unicode"
)

func Tokenize(input string) []string {
	var tokens []string
	var currentToken string
	for _, char := range input {
		if unicode.IsSpace(char) || char == '"' {
			continue
		}
		if char == '{' || char == '}' || char == '[' || char == ']' || char == ',' || char == ':' {
			if currentToken != "" {
				tokens = append(tokens, currentToken)
				currentToken = ""
			}
			tokens = append(tokens, string(char))
		} else {
			currentToken += string(char)
		}
	}
	if currentToken != "" {
		tokens = append(tokens, currentToken)
	}
	return tokens
}

func isNumber(token string) bool {
	for _, char := range token {
		if !unicode.IsDigit(char) {
			return false
		}
	}
	return true
}

func Parse(tokens []string) (interface{}, []string) {
	if len(tokens) == 0 {
		return nil, tokens
	}
	token := tokens[0]
	switch token {
	case "{":
		return parseObject(tokens[1:])
	case "[":
		return parseArray(tokens[1:])
	case "true", "false":
		return token == "true", tokens[1:]
	case "null":
		return nil, tokens[1:]
	default:
		if isNumber(token) {
			number, _ := strconv.Atoi(token)
			return number, tokens[1:]
		}
		return token, tokens[1:]
	}
}

func parseObject(tokens []string) (map[string]interface{}, []string) {
	obj := make(map[string]interface{})
	for len(tokens) > 0 {
		if tokens[0] == "}" {
			return obj, tokens[1:]
		}
		key := tokens[0]
		tokens = tokens[1:]
		if tokens[0] != ":" {
			return nil, tokens
		}
		tokens = tokens[1:]
		var value interface{}
		value, tokens = Parse(tokens)
		obj[key] = value
		if tokens[0] == "," {
			tokens = tokens[1:]
		} else if tokens[0] == "}" {
			return obj, tokens[1:]
		} else {
			return nil, tokens
		}
	}
	return nil, tokens
}

func parseArray(tokens []string) ([]interface{}, []string) {
	var arr []interface{}
	for len(tokens) > 0 {
		if tokens[0] == "]" {
			return arr, tokens[1:]
		}
		var value interface{}
		value, tokens = Parse(tokens)
		arr = append(arr, value)
		if tokens[0] == "," {
			tokens = tokens[1:]
		} else if tokens[0] == "]" {
			return arr, tokens[1:]
		} else {
			return nil, tokens
		}
	}
	return nil, tokens
}
