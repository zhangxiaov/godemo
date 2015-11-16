package main

import (
	"html/template"
	"log"
	"net/http"
)

func SayHello(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello"))
}

func loginHandler(w http.ResponseWriter, req *http.Request) {
	t, err := template.ParseFiles("template/html/login.html")
	if err != nil {
		log.Println(err)
	}
	t.Execute(w, nil)
}

func main() {
	http.HandleFunc("/hello", SayHello)
	http.HandleFunc("/login", loginHandler)
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
