package main

import (
	"encoding/json"
	"log"
	"net/http"

	"embeddednews/internal/news"
)

func main() {
	http.HandleFunc("/news", func(w http.ResponseWriter, r *http.Request) {
		items, err := news.Fetch()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		cleaned := make([]news.Item, len(items))
		for i, it := range items {
			cleaned[i].Title = news.CleanText(it.Title)
			cleaned[i].Link = it.Link
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(cleaned)
	})

	http.HandleFunc("/article", func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.Query().Get("url")
		if url == "" {
			http.Error(w, "missing url", http.StatusBadRequest)
			return
		}
		text, err := news.FetchArticleText(url)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Write([]byte(text))
	})

	log.Println("serving on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
