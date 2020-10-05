package main

import (
	"fmt"
	"os"
	"strings"
	"flag"
	"math/rand"
	"unicode"
)

func dumbmeme(s string){
	fmt.Println(s)
	for i, v := range s[1:]{
		fmt.Printf("%c", v)
		for j := 0; j < i; j++{
			fmt.Printf("%c", ' ')
		}
		fmt.Printf("%c\n", v)
	}
}

func stupidmeme(s string){
	r := rand.New(rand.NewSource(99))
	for _, v := range s{
		if(r.Int() % 2 == 0){
			fmt.Printf("%c", v)
		}else{
			fmt.Printf("%c", unicode.ToUpper(v))
		}
	}
	fmt.Println("")
}

func main(){
	boolFlag := flag.Bool("s", false, "If this is set then the args will be printed in a stupid manner")
	flag.Parse()
	if len(flag.Args()) < 1{
		println("This program expects at least one non-flag argument")
		os.Exit(-1)
	}
	concatenated_args := strings.Join(flag.Args(), " ")
	if *boolFlag{
		stupidmeme(concatenated_args)
	}else{
		dumbmeme(concatenated_args)
	}
	os.Exit(0)
}
