package main

import (
	"net/http"
	"fmt"
)

func main() {
	fmt.Println("Listening on 8080 ")
	http.ListenAndServe(":8080", ChiRouter().InitRouter())
}


