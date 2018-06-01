package main

import (
	"fmt"
	"net/http"
	"os"

	"../handlers"
)

func main() {
	var port string
	if len(os.Args) >= 2 {
		port = ":" + os.Args[1]
	} else {
		port = ":8000"
	}
	http.HandleFunc("/items/", handlers.HandlerItem)
	http.HandleFunc("/items", handlers.HandlerItemsList)
	http.HandleFunc("/", func(arg1 http.ResponseWriter, arg2 *http.Request) {
		fmt.Println("default")
	})
	if err := http.ListenAndServe(port, nil); err != nil {
		panic(err)
	}
}
