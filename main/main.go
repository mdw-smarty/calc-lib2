package main

import (
	"log"
	"os"

	"github.com/mdw-smarty/calc-lib2"
	"github.com/mdw-smarty/calc-lib2/handlers"
)

func main() {
	calculator := &calc.Addition{}
	handler := handlers.NewHandler(calculator, os.Stdout)
	err := handler.Handle(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
}
