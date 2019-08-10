
/*
Parsing a json file and creating html templates 

1. Get json file and Parse json
2. Create handler
3. Create html templates

*/

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

var story Story
var tpl *template.Template

func init() {
	// parse html template
	tpl = template.Must(template.ParseFiles("index.html"))

	// parse JSON file into story struct 
	// can also use json.NewDecoder and all that 
	jsonFile, err := os.Open("gopher.json")
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return
	}

	defer jsonFile.Close()

	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Error reading JSON data:", err)
		return
	}

	json.Unmarshal(jsonData, &story)

    // for _, chapter := range story {
	// 	fmt.Println(chapter)
	// }

}

func main() {
	r := mux.NewRouter()
	
	r.HandleFunc("/{chapter}", displayPage).Methods("GET")

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8080",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())

}

type Story map[string]Chapter 

type Chapter struct {
	Title   string `json:"title"`
	Paragraphs []string `json:"story"`
	Options []Option `json:"options"`
}

type Option struct {
	Text string `json:"text"`
	Chapter string `json:"arc"`
}


func displayPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	// fmt.Fprintln(w, story[vars["chapter"]])
	// f := req.FormValue("first")
	// l := req.FormValue("last")
	// s := req.FormValue("subscribe") == "on"

	err := tpl.ExecuteTemplate(w, "index.html", story[vars["chapter"]])
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Fatalln(err)
	}
}

