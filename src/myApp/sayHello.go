package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

const SERVRE_PORT = 8081
const SERVER_DOMAN = "localhost"

func rootHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello Go !")
}

func main() {
	http.HandleFunc("/", rootHandler)

	err := http.ListenAndServe(
		fmt.Sprintf(":%d", SERVRE_PORT),
		nil)

	//err := http.ListenAndServeTLS(
	//	fmt.Sprintf(":%d", SERVRE_PORT),
	//	"rui.crt",
	//	"rui.key",
	//	nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
