package main

import (
	"fmt"
	"net/http"
	"rsc.io/quote"
)


func home (w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is my home page"))
}

func recipe (w http.ResponseWriter, r *http.Request) {

    userId := r.PathValue("Id")

    if userId != "" {
        w.Write([]byte("userId " + userId))
    } else {
        w.Write([]byte("This is my recipe page"))
    }
}

func main() {
    router := http.NewServeMux()
	router.HandleFunc("GET /", home)
    router.HandleFunc("GET /recipes", recipe)
    router.HandleFunc("GET /recipes/{Id}", recipe)
    
	fmt.Println(quote.Go())
    http.ListenAndServe(":8080", router)
}
