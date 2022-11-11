package main

import "fmt"

//* ABSTRACT FACTORY IMPLEMENTATION EXAMPLE *//

type INotificationFactory interface {
	SendNotification()
	GetSender() ISender
}

type ISender interface {
	GetSenderMethod() string
	GetSenderChannel() string
}

//* SMS notifications

type SMSNotification struct {
}

func (SMSNotification) SendNotification() {
	fmt.Println("Sending notification via SMS")
}

func (SMSNotification) GetSender() ISender {
	return SMSNotificationSender{}
}

type SMSNotificationSender struct {
}

func (SMSNotificationSender) GetSenderMethod() string {
	return "SMS"
}

func (SMSNotificationSender) GetSenderChannel() string {
	return "Twilio"
}

//* Email notifications

type EmailNotification struct {
}

func (EmailNotification) SendNotification() {
	fmt.Println("Sending notification via Email")
}

func (EmailNotification) GetSender() ISender {
	return EmailNotificationSender{}
}

type EmailNotificationSender struct {
}

func (EmailNotificationSender) GetSenderMethod() string {
	return "Email"
}

func (EmailNotificationSender) GetSenderChannel() string {
	return "SendGrid"
}

//* Factory function
func getNotificationFactory(notificationType string) (INotificationFactory, error) {
	switch notificationType {
	case "SMS":
		return SMSNotification{}, nil
	case "Email":
		return EmailNotification{}, nil
	default:
		return nil, fmt.Errorf("notification type %s not recognized", notificationType)
	}
}

func main() {
	notificationFactory, _ := getNotificationFactory("SMS")
	fmt.Println(notificationFactory.GetSender().GetSenderMethod(), notificationFactory.GetSender().GetSenderChannel())

	// Error example
	notificationFactory, err := getNotificationFactory("Push")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(notificationFactory.GetSender().GetSenderMethod(), notificationFactory.GetSender().GetSenderChannel())
	}
}
