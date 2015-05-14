package main

import (
	"fmt"
	"msg_queue"
	"net/http"
)

/*
type GreetingHandler struct{}

func (router GreetingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Greetings!")
}
*/

func GreetingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Greetings")
	msgChnl <- msg_queue.newMsg(msg_queue.GREETING, "", "", w)
}
