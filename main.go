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
	for i, v := range concatenated_args{
		if i == 0{
			fmt.Println(concatenated_args)
			continue
		}
		fmt.Printf("%c", v)
		for j := 1; j < i; j++{
			fmt.Printf("%c", ' ')
		}
		fmt.Printf("%c\n", v)
	}
	os.Exit(0)
}
