package server

import (
	"log"
	"net/http"
)

func CreateServer(address string, db string) {
	http.HandleFunc("/user", GetUserHandler(db))
	http.HandleFunc("/product", GetProductHandler(db))

	err := http.ListenAndServe(address, nil)
	if err != nil {
		log.Fatal(err)
	}
}
