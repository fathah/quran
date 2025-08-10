package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	// Read port and api-secret from env variables
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	apiSecret := os.Getenv("API_SECRET")
	if apiSecret == "" {
		log.Fatal("API_SECRET environment variable not set")
	}

	imageDir := "./images"

	http.HandleFunc("/page/", func(w http.ResponseWriter, r *http.Request) {
		// Check api-secret header
		if r.Header.Get("x-api-secret") != apiSecret {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		pageStr := r.URL.Path[len("/page/"):]
		pageNum, err := strconv.Atoi(pageStr)
		if err != nil || pageNum < 0 || pageNum > 604 {
			http.NotFound(w, r)
			return
		}

		// Compose image path, only .png
		imagePath := filepath.Join(imageDir, fmt.Sprintf("%d.png", pageNum))

		if _, err := os.Stat(imagePath); os.IsNotExist(err) {
			http.NotFound(w, r)
			return
		}

		// Serve PNG image
		w.Header().Set("Content-Type", "image/png")
		http.ServeFile(w, r, imagePath)
	})

	log.Printf("Server starting on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
