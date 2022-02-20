package fcm_client

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
)

var ctx context.Context

func GetFcmConnection() *messaging.Client {

	// There are different ways to add credentials on init.
	// if we have a path to the JSON credentials file, we use the GOOGLE_APPLICATION_CREDENTIALS env var
	//err := os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", c.Firebase.Credentials)
	//if err != nil {
	//	return
	//}
	// or pass the file path directly
	opts := []option.ClientOption{option.WithCredentialsFile("/home/wise/Documents/web/cljm/internal/datasources/fcm_client/google-services.json")}
	// opts := []option.ClientOption{option.WithCredentialsFile("./datasources/fcm_client/google-services.json")}
	//
	// if we have a raw JSON credentials value, we use the FIREBASE_CONFIG env var
	// err = os.Setenv("FIREBASE_CONFIG", "{...}")
	// if err != nil {
	// 	return
	// }

	// or we can pass the raw JSON value directly as an option
	//opts := []option.ClientOption{option.WithAPIKey("AIzaSyARI2UaZJEJBFjTVp3RfrYr0P0nAEpSa6s")}
	//cfg := firebase.Config{
	//	AuthOverride:     nil,
	//	DatabaseURL:      "",
	//	ProjectID:        "itara-89760",
	//	ServiceAccountID: "",
	//	StorageBucket:    "",
	//}
	app, err := firebase.NewApp(ctx, nil, opts...)
	if err != nil {
		log.Fatalf("new firebase app: %s", err)
	}

	fcmClient, err := app.Messaging(context.TODO())
	if err != nil {
		log.Fatalf("messaging: %s", err)
	}
	log.Println("fcm connection successful")
	return fcmClient

}
