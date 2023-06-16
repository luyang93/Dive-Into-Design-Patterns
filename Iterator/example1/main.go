package main

import "fmt"

type Profile struct {
	Email string
}

type ProfileIterator interface {
	GetNext() *Profile
	HasMore() bool
}

type SocialNetwork interface {
	CreateFriendsIterator(profileId string) ProfileIterator
	CreateCoworkersIterator(profileId string) ProfileIterator
}

type WeChat struct{}

type WeChatIterator struct {
	weChat    *WeChat
	profileId string
	typeStr   string
	position  int
	cache     []*Profile
}

func (w *WeChatIterator) GetNext() *Profile {
	if w.HasMore() {
		profile := w.cache[w.position]
		w.position++
		return profile
	}
	return nil
}

func (w *WeChatIterator) HasMore() bool {
	if w.cache == nil {
		w.cache = w.weChat.SocialGraphRequest(w.profileId, w.typeStr)
	}
	return w.position < len(w.cache)
}

func (w *WeChat) CreateFriendsIterator(profileId string) ProfileIterator {
	return &WeChatIterator{weChat: w, profileId: profileId, typeStr: "friends"}
}

func (w *WeChat) CreateCoworkersIterator(profileId string) ProfileIterator {
	return &WeChatIterator{weChat: w, profileId: profileId, typeStr: "coworkers"}
}

func (w *WeChat) SocialGraphRequest(profileId, typeStr string) []*Profile {
	// Here, you should implement your own logic to get the data.
	// The returned data is a slice of pointers to Profile.
	return []*Profile{}
}

type SocialSpammer struct{}

func (s *SocialSpammer) Send(iterator ProfileIterator, message string) {
	for iterator.HasMore() {
		profile := iterator.GetNext()
		fmt.Println("Sending email to ", profile.Email, " with message: ", message)
	}
}

type Application struct {
	network SocialNetwork
	spammer *SocialSpammer
}

func (a *Application) Config() {
	// Here we are using WeChat as the example.
	// You should implement the logic to decide which network to use.
	a.network = &WeChat{}
	a.spammer = &SocialSpammer{}
}

func (a *Application) SendSpamToFriends(profileId string) {
	iterator := a.network.CreateFriendsIterator(profileId)
	a.spammer.Send(iterator, "Very important message")
}

func (a *Application) SendSpamToCoworkers(profileId string) {
	iterator := a.network.CreateCoworkersIterator(profileId)
	a.spammer.Send(iterator, "Very important message")
}

func main() {
	app := &Application{}
	app.Config()
	app.SendSpamToFriends("profileId1")
	app.SendSpamToCoworkers("profileId2")
}
