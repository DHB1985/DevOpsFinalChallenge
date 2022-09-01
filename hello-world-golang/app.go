package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	//"flag"
	"os"
)
var port = os.Getenv("PORT")
var scores = make(map[string]int)
// var port = ":"+strconv.Itoa(3002)
func main() {

	
	// //port := 3002
    if port == "" {
        port = ":"+strconv.Itoa(3002)
    }

	http.HandleFunc("/hello", HelloServer)
	http.HandleFunc("/inc-score", IncrementCounter)
	http.HandleFunc("/get-scores", GetScores)
	http.ListenAndServe(port, nil)

}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	//var port1 = port + "hola"
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
	//fmt.Fprintf(w, "Hello, %s!", port1)

}

// IncrementCounter increments some "score" for a user
func IncrementCounter(w http.ResponseWriter, r *http.Request) {
	name, ok := r.URL.Query()["name"]
	if !ok {
		w.WriteHeader(http.StatusOK)
	}
	scores[name[0]] += 1
	w.WriteHeader(http.StatusOK)
}

// GetScores gets all the scores for all users
func GetScores(w http.ResponseWriter, r *http.Request) {
	b, _ := json.Marshal(scores)
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
