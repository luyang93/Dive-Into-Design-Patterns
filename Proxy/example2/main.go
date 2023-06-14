package main

import (
	"fmt"
	"sync"
)

type RemoteImage struct {
	url string
}

func NewRemoteImage(url string) *RemoteImage {
	return &RemoteImage{url: url}
}

func (ri *RemoteImage) displayImage() {
	fmt.Println("Displaying image from URL:", ri.url)
	// Code to load and display the image
}

type ImageProxy struct {
	url   string
	image *RemoteImage
	mu    sync.Mutex
}

func NewImageProxy(url string) *ImageProxy {
	return &ImageProxy{
		url:   url,
		image: nil,
	}
}

func (ip *ImageProxy) displayImage() {
	ip.mu.Lock()
	defer ip.mu.Unlock()

	if ip.image == nil {
		fmt.Println("New image from URL:", ip.url)
		ip.image = NewRemoteImage(ip.url)
	}
	ip.image.displayImage()
}

func main() {
	imageProxy := NewImageProxy("http://example.com/image.jpg")
	imageProxy.displayImage() // Image is loaded from the server
	imageProxy.displayImage() // Image is loaded from the cache
}
