package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	metrics = `# HELP covid_num Number of covid confirmed cases.
# TYPE covid_num counter
covid_num{region="sg", t="total"} 87892
covid_num{region="cn", t="total"} 116099
covid_num{region="sg", t="current"} 13162
covid_num{region="cn", t="current"} 2720`
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}

func ShowMetrics(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, metrics)
}

func main() {
	http.HandleFunc("/", sayHello)
	http.HandleFunc("/metrics", ShowMetrics)
	log.Println("server start...")
	if err := http.ListenAndServe("0.0.0.0:80", nil); err != nil {
		log.Fatal("server start err:" + err.Error())
	}
}
