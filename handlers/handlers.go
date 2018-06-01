package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"../esDriver"
)

func HandlerItem(w http.ResponseWriter, req *http.Request) {
	if req.Header.Get("Content-Type") == "application/json" {
		fmt.Println("esta pidienod un json")

	}
	id := strings.TrimPrefix(req.URL.Path, "/items/")
	if len(id) > 0 {
		switch req.Method {
		case "GET":
			data, err := esDriver.GetDocFromIndex("items", id)
			if err == nil {
				json.NewEncoder(w).Encode(data)
				// w.Write([]byte(data))
			} else {
				fmt.Println(*err)
			}
		case "DELETE":
			data, err := esDriver.DeleteDoc("items", id)
			if err == nil {
				fmt.Println(data)
				json.NewEncoder(w).Encode(data)
				// w.Write([]byte(*data))
			} else {
				fmt.Println(*err)
			}
		}
	}
}
func HandlerItemsList(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Header.Get("User-Agent"))
	if req.Header.Get("Content-Type") == "application/json" {
		fmt.Println("esta pidienod un json")
	}
	switch req.Method {
	case "POST":
		fmt.Println("request POST")
		w.Write([]byte("{'saludo':'es un POST'}"))
	case "GET":
		fmt.Println("request GET")
		data, err := esDriver.GetDocs("items")
		if err == nil {
			json.NewEncoder(w).Encode(data.Hits.Hits)
			// w.Write([]byte(data))
		} else {
			fmt.Println("erro")
		}
	case "DELETE":
		fmt.Println("request DELETE")
	case "PUT":
		fmt.Println("request PUT")
	case "OPTIONS":
		fmt.Println("request OPTIONS")
	}
	// }
}
func copyHeader(wHeader, respHeader http.Header) {
	for k, vv := range respHeader {
		for _, v := range vv {
			wHeader.Add(k, v)
		}
	}
}
