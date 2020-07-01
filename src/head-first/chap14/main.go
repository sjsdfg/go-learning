package main

import (
	"log"
	"net/http"
)

func write(writer http.ResponseWriter, msg string) {
	_, err := writer.Write([]byte(msg))
	if err != nil {
		log.Fatal(err)
	}
}

func englishHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "english")
}

func chineseHandler(writer http.ResponseWriter, request *http.Request) {
	write(writer, "chinese")
}

func main() {
	http.HandleFunc("/english", englishHandler)
	http.HandleFunc("/chinese", chineseHandler)

	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
