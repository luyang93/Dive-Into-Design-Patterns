package main

import "fmt"

type Observer interface {
	Update(message string)
}

type Reader struct {
	Name string
}

func (r *Reader) Update(message string) {
	fmt.Printf("Notification to %s: %s\n", r.Name, message)
}

type Subject interface {
	RegisterObserver(observer Observer)
	RemoveObserver(observer Observer)
	NotifyObservers()
}

type NewsletterPublisher struct {
	Readers          []Observer
	LatestNewsletter string
}

func (np *NewsletterPublisher) RegisterObserver(observer Observer) {
	np.Readers = append(np.Readers, observer)
}

func (np *NewsletterPublisher) RemoveObserver(observer Observer) {
	for i, ob := range np.Readers {
		if ob == observer {
			np.Readers = append(np.Readers[:i], np.Readers[i+1:]...)
			break
		}
	}
}

func (np *NewsletterPublisher) Publish(newsletter string) {
	np.LatestNewsletter = newsletter
	np.NotifyObservers()
}

func (np *NewsletterPublisher) NotifyObservers() {
	for _, observer := range np.Readers {
		observer.Update(np.LatestNewsletter)
	}
}

func main() {
	publisher := &NewsletterPublisher{}

	alice := &Reader{"Alice"}
	bob := &Reader{"Bob"}
	charlie := &Reader{"Charlie"}

	publisher.RegisterObserver(alice)
	publisher.RegisterObserver(bob)

	publisher.Publish("New edition of the newsletter is out!")

	publisher.RegisterObserver(charlie)
	publisher.RemoveObserver(alice)

	publisher.Publish("Another edition of the newsletter is here!")
}
