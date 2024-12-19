package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func scan(stroka string) []string {
	wordsDig := []string{}
	symbols := []string{}
	closed := true
	sls := []string{}
	for i, v := range stroka {
		if string(v) == "\"" {
			closed = !closed
			if closed {
				word := strings.Join(sls, "")
				wordsDig = append(wordsDig, word)
				sls = []string{}
			}
		} else if !closed {
			sls = append(sls, string(v))
		} else {
			if string(v) != " " {
				symbols = append(symbols, string(v))
				if i == 0 && len(wordsDig) < len(symbols) {
					panic(errors.New("Сначала строка!"))
				}
			}
		}
	}
	if len(symbols) > 1 {
		if len(symbols) > 2 {
			symbols = []string{symbols[0], strings.Join(symbols[1:], "")}
		}
		symbols[0], symbols[1] = symbols[1], symbols[0]
		symbols[0] = symbols[0] + "$"
	}
	if len(wordsDig) == 2 && (len(wordsDig[0]) > 10 || len(wordsDig[1]) > 10) {
		panic(errors.New("Длина одного или двух слов более 10-ти символов"))
	}
	if len(wordsDig) == 1 && len(wordsDig[0]) > 10 {
		panic(errors.New("Длина слова более 10-ти символов"))
	}
	if len(symbols) == 2 && !isStringInteger(strings.Replace(symbols[0], "$", "", 1)) {
		panic(errors.New("Можно работать только с целым числом"))
	}
	if len(symbols) == 2 {
		val, _ := strconv.Atoi(strings.Replace(symbols[0], "$", "", 1))
		if val < 1 || val > 10 {
			panic(errors.New("Число должно быть в диапазоне от 1 до 10."))
		}
	}
	wordsDig = append(wordsDig, symbols...)
	return wordsDig
}

func isStringInteger(s string) bool {
	if s == "" {
		return false
	}
	_, err := strconv.Atoi(s)
	return err == nil
}

func dublLine(s string, val int) (result string) {
	s_ := s
	for i := 1; i < val; i++ {
		s += s_
	}
	return s
}

func splitString(s string, val int) (result string) {
	val = len(s) / val
	for i := 0; i < val; i++ {
		result += string(s[i])
	}
	return
}

func trim(s string) string {
	if len(s) > 40 {
		s = s[:40] + "..."
	}
	return s
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	res := scan(scanner.Text())
	val_1, val_2, symb := res[0], res[1], res[2]
	val_2_int, _ := strconv.Atoi(strings.Replace(val_2, "$", "", 1))
	if symb == "+" && string(val_2[len(val_2)-1]) != "$" {
		fmt.Println("\"" + val_1 + val_2 + "\"")
	} else if symb == "-" && string(val_2[len(val_2)-1]) != "$" {
		fmt.Println("\"" + strings.ReplaceAll(val_1, val_2, "") + "\"")
	} else if symb == "*" && string(val_2[len(val_2)-1]) == "$" {
		fmt.Println("\"" + trim(dublLine(val_1, val_2_int)) + "\"")
	} else if symb == "/" && string(val_2[len(val_2)-1]) == "$" {
		fmt.Println("\"" + splitString(val_1, val_2_int) + "\"")
	} else {
		panic(errors.New("Неподдерживаемая операция!"))
	}
}
