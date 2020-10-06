package main

import (
	"fmt"
	"os"
	"strings"
	"flag"
	"math/rand"
	"unicode"
	"io/ioutil"
	"net"
	"strconv"
	"net/http"
	"log"
)

func dumbmeme(s string) string{
	var b strings.Builder
	s = strings.ToUpper(s)
	b.WriteString(strings.ReplaceAll(s, "\n", " "))
	b.WriteString("\n")
	for i, v := range s[1:]{
		fmt.Fprintf(&b, "%c", v)
		for j := 0; j < i; j++{
			fmt.Fprintf(&b, "%c", ' ')
		}
		fmt.Fprintf(&b, "%c\n", v)
	}
	return b.String()
}

func stupidmeme(s string) string{
	r := rand.New(rand.NewSource(99))
	var b strings.Builder
	for _, v := range s{
		if(r.Int() % 2 == 0){
			fmt.Fprintf(&b, "%c", v)
		}else{
			fmt.Fprintf(&b, "%c", unicode.ToUpper(v))
		}
	}
	b.WriteString("\n")
	return b.String()
}

func handleClientRequest(conn net.Conn, boolFlag *bool){
	buf := make([]byte, 1024)
	var err error = nil
	var n int
	for err == nil{
		n, err = conn.Read(buf)
		if err != nil{
			continue
		}
		s := string(buf[:n])
		if *boolFlag{
			conn.Write([]byte(stupidmeme(s)))
		}else{
			conn.Write([]byte(dumbmeme(s)))
		}
	}
	conn.Close()
}

func indexHandler(w http.ResponseWriter, r *http.Request){
	http.ServeFile(w, r, "index.html")
}

func inputHandler(w http.ResponseWriter, r *http.Request){
	if err := r.ParseForm(); err != nil{
		println("Error reading query")
	}
	if r.FormValue("stupid") == "y"{
		fmt.Fprintf(w, "%s", stupidmeme(r.FormValue("input")))
	}else{
		fmt.Fprintf(w, "%s", dumbmeme(r.FormValue("input")))
	}
}

func httpServer(){
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/memestring", inputHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func server(boolFlag *bool, portFlag *int, hostFlag *string){
	var l net.Listener
	var err error
	fmt.Println(*hostFlag + ":" + strconv.Itoa(*portFlag));
	if l, err = net.Listen("tcp", *hostFlag + ":" + strconv.Itoa(*portFlag)); err != nil{
		println("Could not listen")
		println(err.Error())
		os.Exit(-1)
	}
	defer l.Close()
	for {
		var conn net.Conn
		if conn, err = l.Accept(); err != nil{
			println("Error accepting a connection")
			continue
		}
		go handleClientRequest(conn, boolFlag)

	}
}

func main(){
	mdFlag := flag.Bool("d", false, "If this is set then the output will be enclosed in ``` for markdown")
	boolFlag := flag.Bool("s", false, "If this is set then the args will be printed in a stupid manner")
	inputFlag := flag.String("i", "" , "Set what the input is (<filename> or - for stdin), by default it's the non-flag arguments")
	listenFlag := flag.Bool("l", false, "Set this as a server listening on port 1234 of localhost by default")
	portFlag := flag.Int("p", 1234, "When used in conjunction with -l, use to specify port, defaults to 1234")
	hostFlag := flag.String("h", "localhost", "When used in conjuction with -l, specifies the IP address to listen to, defaults to localhost")
	httpFlag := flag.Bool("http", false, "Set this to run it as an http server")
	flag.Parse()

	if *listenFlag{
		server(boolFlag, portFlag, hostFlag)
		os.Exit(0)
	}else if *httpFlag{
		httpServer()
		os.Exit(0)
	}
	var s string
	if *inputFlag != ""{
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
		fmt.Print(stupidmeme(s))
	}else{
		fmt.Print(dumbmeme(s))
	}

	if *mdFlag{
		fmt.Println("```")
	}
	os.Exit(0)
}
