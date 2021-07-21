package main

import (
	"encoding/json"
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"net/http"

	"github.com/nfnt/resize"
)

func main() {
	http.HandleFunc("/", ImageResize)
	log.Println("Listening on 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func handleImageErr(w http.ResponseWriter, err string) {
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprint(w, err)
	return
}

func ImageResize(w http.ResponseWriter, r *http.Request) {

	var req struct {
		Image string `json:"image"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		handleImageErr(w, "No image url supplied")
		return
	}

	if req.Image == "" {
		handleImageErr(w, "No image url supplied")
		return
	}

	resp, err := http.Get(req.Image)
	if err != nil {
		handleImageErr(w, "Invalid url supplied")
		return
	}
	defer resp.Body.Close()

	m, _, err := image.Decode(resp.Body)
	if err != nil {
		handleImageErr(w, "Failed decoding image")
		return
	}

	newImage := resize.Resize(160, 0, m, resize.Lanczos3)
	err = jpeg.Encode(w, newImage, nil)
	if err != nil {
		handleImageErr(w, "Failed encoding re-sized image")
	}
}
