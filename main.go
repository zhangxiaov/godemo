package main

import (
	// "fmt"
	"html/template"
	"log"
	"net/http"
	// "strings"
)

func SayHello(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello"))
}

func homeHandler(w http.ResponseWriter, req *http.Request) {
	t, err := template.ParseFiles("template/html/home.html")
	if err != nil {
		log.Println(err)
	}
	t.Execute(w, nil)
}

func postHandler(w http.ResponseWriter, req *http.Request) {

	req.ParseForm()

	if req.Form["postTitle"] != nil && len(req.Form["postTitle"][0]) > 0 {
		log.Println(len(req.Form["postTitle"]))
		urlStr := "/home"
		http.Redirect(w, req, urlStr, http.StatusFound)
	}

	t, err := template.ParseFiles("template/html/post.html")
	if err != nil {
		log.Println(err)
	}
	t.Execute(w, nil)
}

func loginHandler(w http.ResponseWriter, req *http.Request) {

	req.ParseForm()

	if formMap := req.Form; len(formMap) > 0 {
		if req.Form["userName"][0] == "zhangxinwei" && req.Form["password"][0] == "123456" {

			urlStr := "/home"
			http.Redirect(w, req, urlStr, http.StatusFound)
		}
	}
	// fmt.Println(req.Form)
	// fmt.Println("path", req.URL.Path)
	// fmt.Println("Scheme", req.URL.Scheme)
	// fmt.Println(req.Form["url_long"])
	// for k, v := range req.Form {
	// 	fmt.Println("key:", k)
	// 	// join() 方法用于把数组中的所有元素放入一个字符串。
	// 	// 元素是通过指定的分隔符进行分隔的
	// 	fmt.Println("val:", strings.Join(v, ""))
	// }

	t, err := template.ParseFiles("template/html/login.html")
	if err != nil {
		log.Println(err)
	}
	t.Execute(w, nil)
}

func main() {
	http.HandleFunc("/hello", SayHello)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/home", homeHandler)
	http.HandleFunc("/post", postHandler)
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
