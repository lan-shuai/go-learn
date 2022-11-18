package main

import (
	. "fmt"
	"html"
	"net/http"
	"time"
)

var count int = 1

func Task(w http.ResponseWriter, r *http.Request) {
	Printf(" ------ here is Task[%d] ------- \n", count)
	defer r.Body.Close()

	// 模拟延时
	time.Sleep(time.Second * 2)

	Fprintf(w, "Response.No.%d: Your request is %q\n\n", count, html.EscapeString(r.URL.Path))
	count++
	answer := `{"status:":"OK"}`
	w.Write([]byte(answer))
}

func Task1(w http.ResponseWriter, r *http.Request) {
	Printf(" ------ here is Task1[%d] ------- \n", count)
	defer r.Body.Close()

	// 模拟延时
	time.Sleep(time.Second * 2)

	Fprintf(w, "Response.No[%d]: Your request is %q\n\n", count, html.EscapeString(r.URL.Path))
	count++
	answer := `{"status:":"OK"}`
	w.Write([]byte(answer))
}

func main() {
	Println("===== This is http server =====")

	http.HandleFunc("/hello/world", Task)

	http.HandleFunc("/hello/world1", Task1)

	err := http.ListenAndServe("127.0.0.1:8090", nil)
	if err != nil {
		Println("ListenAndServe error! err is ", err.Error())
	}
}
