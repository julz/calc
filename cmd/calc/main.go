package main

import (
	"fmt"
	"os"

	"github.com/julz/calc"
)

func main() {
	fmt.Println(calc.Must(calc.Eval(os.Args[1])))
}
