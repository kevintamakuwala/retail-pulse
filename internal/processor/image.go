package processor

import (
	"io"
	"math/rand"
	"net/http"
	"time"
)

// ProcessImage downloads and processes an image
func ProcessImage(imageURL string) error {
	// Download image
	resp, err := http.Get(imageURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Read image data
	_, err = io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Simulate processing time (0.1 to 0.4 seconds as per requirements)
	sleepTime := 100 + rand.Intn(300)
	time.Sleep(time.Duration(sleepTime) * time.Millisecond)

	return nil
}
