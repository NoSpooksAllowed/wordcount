package main

import (
	"fmt"
	"os"

	"github.com/NoSpooksAllowed/wordcount/counter"
	"github.com/NoSpooksAllowed/wordcount/input"
)

func main() {
	str := input.ReadInputFromCmd(os.Args)

	fmt.Println(counter.CountWords(str))
}
