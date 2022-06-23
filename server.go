package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

//type Demo int

//func (d Demo) ServeHTTP(w http.ResponseWriter, request *http.Request) {
//
//	io.WriteString(w, "hello shan! i am replying from server side!")
//
//}

func main() {
	http.HandleFunc("/hello", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(res, "hi, %q", html.EscapeString(req.URL.Path))
	})

	////fmt.Println("server started...")
	////log.Fatal(http.ListenAndServe(":8080", nil))
	//
	//http.HandleFunc("/ping", func(res http.ResponseWriter, req *http.Request) {
	//	fmt.Fprintf(res, "hello shani, %q", html.EscapeString(req.URL.Path))
	//})
	//var d Demo
	//
	//mux := http.NewServeMux()
	//mux.Handle("/shani/", d)

	fmt.Println("server started...")
	log.Fatal(http.ListenAndServe(":8060", nil))
}

//handler: is a interface that have method signature of serveHTTP(ResponseWriter,*Request)
//implementing this interface means any method which signature serveHTTP(ResponseWriter, *Request)
//http.ResponseWriter: is an interface which have method signature :Write([]byte)(int,error),io.Writer,
//io.WriteString
//http.*Request:its an struct  holds the data from the client
//

//the server multiplexer
//http.NewServeMux returns ServeMux

//type ServeMux struct {
//	mu    sync.RWMutex
//	m     map[string]int
//	es    []muxEntry
//	hosts bool
//}

//it's an abstraction of all routes and handlers

//ServerMux have 4 public methods

//func(mux *ServeMux) Handle(pattern string,handler http.Handler)
//HandleFunc(pattern string, handler func(http ResponseWriter,*http.request))
//Handler(r *Request)(h handler, patter string)
//ServeHTTP(w http.ResponseWriter, r *http.Request)

//handle takes a type which implements the handler interface while HandleFunc takes standalone function

//HandleFunc :http.HandleFunc takes a standalone function instead of taking a type which implements Handler interface.
