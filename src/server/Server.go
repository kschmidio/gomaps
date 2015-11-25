package main

import (
	"encoding/json"
	"net/http"
	"fmt"
)

type Message struct {
	Text string
}

type Job struct {
	Title string
	City  string
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/about/", about)
	http.HandleFunc("/api/jobs/", jobs)
	http.ListenAndServe(":8080", nil)
}

func about(w http.ResponseWriter, r *http.Request) {
	m := Message{"Welcome to Rest Services implemented with GO"}
	b, err := json.Marshal(m)

	checkError(err)

	w.Write(b)
}

func jobs(w http.ResponseWriter, r *http.Request) {
	keyword := r.FormValue("keyword")
	company := r.FormValue("company")
	xml := getRssFeed(keyword, company)
	rss := parseXml(xml)
	jobDTOs := calculateJobsPerCity(rss.Channel)
	
	fmt.Println("found jobs ", len(jobDTOs))
	
	b, err := json.Marshal(jobDTOs)

	checkError(err)

	w.Write(b)
	
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
