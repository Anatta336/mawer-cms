package main

import (
	"fmt"
	"log"
	"mawer/config"
	"mawer/document"
	"net/http"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error loading config:", err)
	}

	prepareRoutes()

	fmt.Println("Starting server on " + cfg.ServerAddress)
	log.Fatal(http.ListenAndServe(cfg.ServerAddress, nil))
}

func prepareRoutes() {
	log.Println("Preparing routes...")

	http.HandleFunc("POST /documents/upload", document.UploadHandler)

	// Static files.
	http.Handle("GET /static/{file}",
		http.StripPrefix("/static/", http.FileServer(http.Dir("public/static"))),
	)

	// Index, with 404 for everything else.
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		log.Println("URL requested:", request.URL.Path)

		if request.URL.Path != "/" {
			http.NotFound(writer, request)
			return
		}

		http.ServeFile(writer, request, "public/index.html")
	})
}
