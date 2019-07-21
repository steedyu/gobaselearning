package main

import (
	"io"
	"log"
	"net/http"
	"html/template"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	hp := `<html>
    <head>
    <title>okkkkkk</title>
    <link rel="stylesheet" href="template/css/main.css" type="text/css" />
    </head>
    <body>
        <h2>this is a test for golang.</h2>
    </body>
    </html>`
	io.WriteString(w, hp)
}

func Hello2(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./template/hello.gohtml")
	if err != nil {
		panic(err)
	}

	data := struct {
		Name string
	}{"John Smith"}

	err = t.Execute(w, data)
	if err != nil {
		panic(err)
	}
}

func StaticServer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/html")
	staticHandler := http.FileServer(http.Dir("./template/"))
	staticHandler.ServeHTTP(w, r)
	return
}


func main() {

	/*
	映射静态文件

	"github.com/julienschmidt/httprouter"
	router.ServeFiles("/page/*filepath", http.Dir("page"))
	也存在相同的功能
	 */
	http.Handle("/template/", http.StripPrefix("/template/", http.FileServer(http.Dir("./template"))))
	http.HandleFunc("/hello", Hello2)
	http.HandleFunc("/", Hello)

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}
