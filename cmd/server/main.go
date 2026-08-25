package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/home", homeHandler)

	fileServer := http.FileServer(http.Dir("./web/static"))

	http.Handle("/static/",
		http.StripPrefix("/static/", fileServer))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./web/js"))))

	fmt.Println("сервер запущен на http://127.0.0.1:8020/home")
	http.ListenAndServe(":8020", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./web/html/index.html")
}
