package services

import (
	"context"
	"firebase.google.com/go/messaging"
	"log"
)

type FcmService interface {
	SendNotification() error
}

type fcmService struct {
	fcm *messaging.Client
}

func NewFcmService(fcm *messaging.Client) FcmService {
	return &fcmService{fcm: fcm}
}

func (f *fcmService) SendNotification() error {
	response, err := f.fcm.Send(context.TODO(), &messaging.Message{
		Data: map[string]string{
			"A nice notification title": "A nice notification title",
		},
		Notification: &messaging.Notification{
			Title: "A nice notification title",
			Body:  "A nice notification body",
		},
		Android: &messaging.AndroidConfig{
			CollapseKey:           "",
			Priority:              "",
			TTL:                   nil,
			RestrictedPackageName: "",
			Data:                  nil,
			Notification: &messaging.AndroidNotification{
				Title:                 "A nice notification title",
				Body:                  "A nice notification body",
				Icon:                  "",
				Color:                 "",
				Sound:                 "",
				Tag:                   "",
				ClickAction:           "",
				BodyLocKey:            "",
				BodyLocArgs:           nil,
				TitleLocKey:           "",
				TitleLocArgs:          nil,
				ChannelID:             "",
				ImageURL:              "",
				Ticker:                "",
				Sticky:                false,
				EventTimestamp:        nil,
				LocalOnly:             false,
				Priority:              0,
				VibrateTimingMillis:   nil,
				DefaultVibrateTimings: false,
				DefaultSound:          false,
				LightSettings:         nil,
				DefaultLightSettings:  false,
				Visibility:            0,
				NotificationCount:     nil,
			},
			FCMOptions: nil,
		},
		Webpush:   nil,
		APNS:      nil,
		Topic:     "properties",
		Condition: "",
	})
	log.Println(response)
	return err

}
