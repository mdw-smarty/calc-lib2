package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/mdw-smarty/calc-lib2"
)

func main() {
	calculator := &calc.Addition{}
	handler := NewHandler(calculator, os.Stdout)
	err := handler.Handle(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
}

type Handler struct {
	calculator *calc.Addition
	stdout     *os.File
}

func NewHandler(calculator *calc.Addition, stdout *os.File) *Handler {
	return &Handler{
		calculator: calculator,
		stdout:     stdout,
	}
}

func (this *Handler) Handle(args []string) error {
	if len(args) != 2 {
		return errors.New("2 operands required")
	}
	a, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	}
	b, err := strconv.Atoi(args[1])
	if err != nil {
		return err
	}
	c := this.calculator.Calculate(a, b)
	_, err = fmt.Fprintf(this.stdout, "%d", c)
	if err != nil {
		return err
	}
	return nil
}
