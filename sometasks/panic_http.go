package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		panic("oh, shit happens, but server will be alive")
	})

	fmt.Println("starting http server.....")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("cannot start server ", err)
	}

}
