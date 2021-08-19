package main

import (
	"complex_calc/calc"
	//"fmt"
	"html/template"
	"log"
	"net/http"
)

type Result struct {
	Expression string
	Answer     string
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

	expr := r.Form.Get("expr")
	result := calc.Calculate(string(expr))
	p := Result{Expression: string(expr), Answer: result}
	t, _ := template.ParseFiles("template/finalpage.html")
	t.Execute(w, p)
	//fmt.Fprintf(w, result)
}
