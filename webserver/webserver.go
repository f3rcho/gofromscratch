package webserver

import (
	"fmt"
	"net/http"
)

func MyWebServer() {
	http.HandleFunc("/", home)
	fmt.Println("Server is listening on port 8080")
	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./webserver/index.html")
}
