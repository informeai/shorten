package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

//ShortLink struct of return links shortening.
type ShortLink struct {
	Short string `json:"shortlink"`
}

//LongUrl struct of parse request url data.
type LongUrl struct {
	Long string `json:"long"`
}

func serveIndex(w http.ResponseWriter, r *http.Request) {
	fBytes, err := ioutil.ReadFile("./templates/index.html")
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	w.Write(fBytes)
}
func shortLinks(w http.ResponseWriter, r *http.Request) {
	var long LongUrl
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&long)
	if err != nil {
		log.Println(err)
		return
	}
	//create shortlink here...
	short := ShortLink{Short: "www.informeai.com.br"}
	jsonBytes, err := json.Marshal(short)
	if err != nil {
		log.Println(err)
		return
	}
	w.Write(jsonBytes)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", serveIndex).Methods("GET")
	router.HandleFunc("/short", shortLinks).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", router))
}
