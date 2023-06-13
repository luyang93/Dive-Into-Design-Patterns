package main

import "fmt"

type Notifier interface {
	Send(message string)
}

type EmailNotifier struct{}

func (e *EmailNotifier) Send(message string) {
	fmt.Println("Sending Email message:", message)
}

type NotifierDecorator struct {
	notifier Notifier
}

func (nd *NotifierDecorator) Send(message string) {
	nd.notifier.Send(message)
}

type WeChatDecorator struct {
	NotifierDecorator
}

func (wd *WeChatDecorator) Send(message string) {
	wd.notifier.Send(message)
	fmt.Println("Sending WeChat message:", message)
}

type MobileDecorator struct {
	NotifierDecorator
}

func (md *MobileDecorator) Send(message string) {
	md.NotifierDecorator.Send(message)
	fmt.Println("Sending Mobile message:", message)
}

type QQDecorator struct {
	NotifierDecorator
}

func (qd *QQDecorator) Send(message string) {
	qd.NotifierDecorator.Send(message)
	fmt.Println("Sending QQ message:", message)
}

// Client
func main() {
	message := "Hello, this is a notification!"

	// Sending notifications using different combinations
	var notifier Notifier = &EmailNotifier{}
	notifier = &WeChatDecorator{NotifierDecorator: NotifierDecorator{notifier}}
	notifier = &MobileDecorator{NotifierDecorator: NotifierDecorator{notifier}}
	notifier = &QQDecorator{NotifierDecorator: NotifierDecorator{notifier}}

	notifier.Send(message)
}
