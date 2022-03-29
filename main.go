package main

import (
	"fmt"
	"net/http"

	"abjs/main/api"
)

func main() {
	fmt.Println("Server is running on port 8080")
	srv := api.NewServer()
	http.ListenAndServe(":8080", srv)
}
