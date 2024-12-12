package main

import (
	"fmt"
	"github.com/bits-and-blooms/bloom/v3" // используем пакет реализующий фильтр Блума
)

func main() {
	filter := bloom.NewWithEstimates(1000000, 0.0001)
	var input string

	fmt.Scan(&input)
	for i := 0; i < len(input); i++ {
		if filter.Test([]byte{input[i]}) { // если символ найден в фильтре - выводим false
			fmt.Println("false")
			return
		} else { //если символ не найден в фильтре, добавляем его в фильтр
			filter.Add([]byte{input[i]})
			if input[i] >= 'a' && input[i] <= 'z' { // если символ является буквой - добавляем ее в фильтр другом регистре
				filter.Add([]byte{input[i] - 32})
			}
			if input[i] >= 'A' && input[i] <= 'Z' {
				filter.Add([]byte{input[i] + 32})
			}
		}
	}
	fmt.Println("true")
}
