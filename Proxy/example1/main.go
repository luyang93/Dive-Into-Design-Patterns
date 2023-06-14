package main

import "fmt"

type ThirdPartyYouTubeLib interface {
	listVideos() string
	getVideoInfo(id string) string
	downloadVideo(id string)
}

type ThirdPartyYouTubeClass struct{}

func (t ThirdPartyYouTubeClass) listVideos() string {
	fmt.Println("Sending an API request to YouTube: listVideos")
	return "Video list"
}

func (t ThirdPartyYouTubeClass) getVideoInfo(id string) string {
	fmt.Printf("Sending an API request to YouTube: getVideoInfo for video with ID %s\n", id)
	return "Video info"
}

func (t ThirdPartyYouTubeClass) downloadVideo(id string) {
	fmt.Printf("Downloading video with ID %s from YouTube\n", id)
}

type CachedYouTubeClass struct {
	service    ThirdPartyYouTubeLib
	listCache  string
	videoCache string
	needReset  bool
}

func (c CachedYouTubeClass) listVideos() string {
	if c.listCache == "" || c.needReset {
		c.listCache = c.service.listVideos()
	}
	return c.listCache
}

func (c CachedYouTubeClass) getVideoInfo(id string) string {
	if c.videoCache == "" || c.needReset {
		c.videoCache = c.service.getVideoInfo(id)
	}
	return c.videoCache
}

func (c CachedYouTubeClass) downloadVideo(id string) {
	if !c.downloadExists(id) || c.needReset {
		c.service.downloadVideo(id)
	}
}

func (c CachedYouTubeClass) downloadExists(id string) bool {
	return false
}

type YouTubeManager struct {
	service ThirdPartyYouTubeLib
}

func (m YouTubeManager) renderVideoPage(id string) {
	info := m.service.getVideoInfo(id)
	fmt.Println("Rendering the video page:", info)
}

func (m YouTubeManager) renderListPanel() {
	list := m.service.listVideos()
	fmt.Println("Rendering the list of video thumbnails:", list)
}

func (m YouTubeManager) reactOnUserInput() {
	m.renderVideoPage("videoID")
	m.renderListPanel()
}

type Application struct{}

func (a Application) init() {
	aYouTubeService := ThirdPartyYouTubeClass{}
	aYouTubeProxy := CachedYouTubeClass{service: aYouTubeService}
	manager := YouTubeManager{service: aYouTubeProxy}
	manager.reactOnUserInput()
}

func main() {
	app := Application{}
	app.init()
}
