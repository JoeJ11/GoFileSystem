package main

const (
	GREETING       = 0
	KV_INSERT      = 1
	KV_DELETE      = 2
	KV_GET         = 3
	KV_UPDATE      = 4
	KVMAN_COUNTKEY = 5
	KVMAN_DUMP     = 6
	KVMAN_SHUTDOWN = 7
)

var (
	msgChnl = make(chan int, 100)
)
