package main

import (
	back "hangman/back"
	"net/http"
)

func main() {
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./css"))))
	http.HandleFunc("/", back.Hangman)
	http.HandleFunc("/hangman", back.Letter)
	http.HandleFunc("/win", back.Win)
	http.HandleFunc("/lose", back.Lose)
	http.HandleFunc("/retry", back.Retry)
	http.ListenAndServe(":8080", nil)
}
