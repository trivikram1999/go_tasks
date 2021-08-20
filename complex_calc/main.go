package main

import (
	"complex_calc/calc"
	//"fmt"
	"html/template"
	"log"
	"net/http"
)

//Result is a struct
type Result struct {
	Expression string
	Answer     interface{}
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/answer", ansPage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./template/index.html")
}

func ansPage(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	var result string
	var err1 error
	expr := r.Form.Get("expr")
	result, err1 = calc.Calculate(string(expr))
	var p Result
	if err1 == nil {
		p = Result{Expression: string(expr), Answer: result}
	} else {
		p = Result{Expression: string(expr), Answer: err1}
	}

	t, _ := template.ParseFiles("template/finalpage.html")
	t.Execute(w, p)
	//fmt.Fprintf(w, result)
}
