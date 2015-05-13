package main

import (
	"fmt"
	"msg_queue"
)

var (
	table = make(map[string]string)
)

func waitForMsg() {
	for true {
		switch <-msg_queue.msgChannel {
		case msg_queue.GREETING:
			{
				w := greeting.getWriter()
				// do nothing
			}
		case msg_queue.KV_INSERT:
			{
				w := handle_insert.getWriter()
				key, val := handle_insert.getNext()
				table[key] = val
			}
		case msg_queue.KV_DELETE:
			{
				w := handle_delete.getWriter()
				key := handle_delete.getNext()
				delete(table, key)
			}
		case msg_queue.KV_GET:
			{
				w := handle_get.getWriter()
				key := handle_get.getNext()
				val := table[key]
				fmt.Fprintln(w, key+" : "+val)
			}
		case msg_queue.KV_UPDATE:
			{
				w := handle_update.getWriter()
				key, val = handle_update()
				table[key] = val
			}
		case msg_queue.KVMAN_COUNTKEY:
			{
				w := handle_counter_key.getWriter()
				fmt.Fprintf(w, "table-size : %d\n", len(table))
			}
		}
	}
}
