package processor

import (
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"net/http"
)

func ProcessImage(imageURL string) (int, error) {
	log.Printf("Downloading image: %s", imageURL)

	// Download image
	resp, err := http.Get(imageURL)
	if err != nil {
		log.Printf("Error downloading image: %s, error: %v", imageURL, err)
		return 0, err
	}
	defer resp.Body.Close()

	// Decode img
	config, _, err := image.DecodeConfig(resp.Body)
	if err != nil {
		log.Printf("Decoding Failed: %v", err)
		return 0, err
	}

	// Calculate perimeter
	width := config.Width
	height := config.Height
	perimeter := 2 * (width + height)
	log.Printf("Image processed: %s, perimeter: %d", imageURL, perimeter)

	return perimeter, err
}
