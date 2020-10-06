package main

import (
	"fmt"
	"os"
	"strings"
	"flag"
	"math/rand"
	"unicode"
	"io/ioutil"
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
	mdFlag := flag.Bool("d", false, "If this is set then the output will be enclosed in ``` for markdown")
	boolFlag := flag.Bool("s", false, "If this is set then the args will be printed in a stupid manner")
	inputFlag := flag.String("i", "args", "Set what the input is (<filename> or - for stdin), by default it's the non-flag arguments")
	flag.Parse()
	var s string
	if *inputFlag != "args"{
		var bytes []byte
		var err error
		if *inputFlag == "-"{
			bytes, err = ioutil.ReadAll(os.Stdin)
			if err != nil{
				println("Error")
				os.Exit(-1)
			}
		}else{
			bytes, err = ioutil.ReadFile(*inputFlag)
			if err != nil{
				println("Error")
				os.Exit(-1)
			}
		}
		s = string(bytes)
	}else if len(flag.Args()) < 1{
		println("This program expects at least one non-flag argument with default input")
		os.Exit(-1)
	}else{
		s = strings.Join(flag.Args(), " ")
	}

	if *mdFlag{
		fmt.Println("```")
	}

	if *boolFlag{
		stupidmeme(s)
	}else{
		dumbmeme(s)
	}

	if *mdFlag{
		fmt.Println("```")
	}
	os.Exit(0)
}
