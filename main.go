package main

import (
	"embed"
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
)

var (
	//go:embed templates
	aaa embed.FS

	//go:embed assets
	bbb embed.FS
)

func main() {
	t := template.Must(template.ParseFS(aaa, "templates/index.gohtml"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_ = t.Execute(w, map[string]any{
			"BG": fmt.Sprintf("/assets/bgs/%d.webp", rand.Intn(5)+1),
		})
	})

	http.HandleFunc("/robots.txt", func(w http.ResponseWriter, r *http.Request) {
		robotsFile, _ := bbb.ReadFile("assets/robots.txt")
		w.Header().Set("Content-Type", "text/plain")
		_, _ = w.Write(robotsFile)
	})
	http.HandleFunc("/sitemap.xml", func(w http.ResponseWriter, r *http.Request) {
		robotsFile, _ := bbb.ReadFile("assets/sitemap.xml")
		w.Header().Set("Content-Type", "text/plain")
		_, _ = w.Write(robotsFile)
	})

	http.Handle("/assets/", http.FileServer(http.FS(bbb)))

	log.Println("server running on port 8080")
	log.Fatalln(http.ListenAndServe(":8080", nil))
}
