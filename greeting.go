package main

import (
	"fmt"
	"net"
	"net/http"
	"bufio"
	"os"
)

/*
type GreetingHandler struct{}

func (router GreetingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Greetings!")
}
*/

func GreetingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handle Greeting")

	conn, err := net.Dial("tcp", "localhost:5000")
	fmt.Println("Connected")

	if (err != nil) {
		fmt.Println("GreetingHandler::DIAL::Error: " + err.Error())
		os.Exit(1)
	}

	fmt.Fprint(conn, "Sending Message\n")

	fmt.Println("Message sent.")
	connBuf := bufio.NewReader(conn)
	str, err := connBuf.ReadString('\n')
	if err != nil {
		fmt.Println("GreetingHandler::READSTRING::Error: " + err.Error())
		os.Exit(1)
	}
	if len(str) > 0 {
		fmt.Println(str)
	}

	fmt.Fprint(w, "Greetings")
	conn.Close()

}
