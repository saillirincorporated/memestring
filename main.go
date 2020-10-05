package main

import (
	"fmt"
	"os"
	"strings"
)

func main(){
	args := os.Args[1:]
	if len(args) < 1{
		println("This program expects exactly one argument")
		os.Exit(-1)
	}
	concatenated_args := strings.Join(args, " ")
	concatenated_args = strings.ToUpper(concatenated_args)
	fmt.Println(concatenated_args)
	for i, v := range concatenated_args[1:]{
		fmt.Printf("%c", v)
		for j := 0; j < i; j++{
			fmt.Printf("%c", ' ')
		}
		fmt.Printf("%c\n", v)
	}
	os.Exit(0)
}
