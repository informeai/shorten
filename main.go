package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/informeai/shorten/base62"
	"github.com/informeai/shorten/entities"
	"github.com/informeai/shorten/store"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
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
	//create rando uint64 and encode link.
	rndUint64 := rand.Uint64()
	shortlink := base62.Encode(rndUint64)

	//create shorten struct
	srt := entities.Shorten{Id: strconv.FormatUint(rndUint64, 10), Url: long.Long, Visits: 0}
	//Insert to database
	db := store.NewStoreSqlite()
	err = db.Insert(srt)
	if err != nil {
		log.Println(err)
		return
	}
	shortlink = fmt.Sprintf("http://localhost:8000/%v", shortlink)
	short := ShortLink{Short: shortlink}
	jsonBytes, err := json.Marshal(short)
	if err != nil {
		log.Println(err)
		return
	}
	w.Write(jsonBytes)
}

func redirectUrl(w http.ResponseWriter, r *http.Request) {
	var shortlink string
	vars := mux.Vars(r)
	shortlink = vars["shortlink"]
	//decode shortlink
	id, err := base62.Decode(shortlink)
	if err != nil {
		log.Println(err)
		return
	}
	//get store and verify if exists shortlink
	db := store.NewStoreSqlite()
	srt, err := db.Get(strconv.FormatUint(id, 10))
	if err != nil {
		log.Println(err)
		return
	}
	//update visits
	err = db.Update(srt)
	if err != nil {
		log.Println(err)
		return
	}
	http.Redirect(w, r, srt.Url, 302)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", serveIndex).Methods("GET")
	router.HandleFunc("/short/create", shortLinks).Methods("POST")
	router.HandleFunc("/{shortlink}", redirectUrl)
	log.Fatal(http.ListenAndServe(":8000", router))
}
