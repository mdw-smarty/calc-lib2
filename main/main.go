package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/mdw-smarty/calc-lib2"
)

func main() {
	calculator := &calc.Addition{}
	a := parseInt(os.Args[1])
	b := parseInt(os.Args[2])
	c := calculator.Calculate(a, b)
	_, err := fmt.Println(c)
	if err != nil {
		panic(err)
	}
}

func parseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
