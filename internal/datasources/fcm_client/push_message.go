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

	opts := []option.ClientOption{option.WithCredentialsJSON([]byte(`{
  "type": "service_account",
  "project_id": "itara-89760",
  "private_key_id": "b9be473eb9c1a5002dab392a0e9e9b6ac25b8537",
  "private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQC9+rIWe62F4HYN\nkR9SKgrQOX5iKtLZWI5/1r6VoUIxnzIEl87/lAE6ec6TDcwa1tT57UEe3AJbKmA6\n1MCgBlwzACS2D0dy0WpdGtf2obo9yKMy6dxEBi2XjttOOWIyyXbKjRfyBfg5ccKO\nUv/hQTXxjX9HCz+zmHipj5bE+8AhwjieFo3440UGcjQ3oAuQdPxSVFzOor5s5Cng\nXttJl5ntTXap3eYS5i0VngHyOnb4Oc26gP7lMwYohHTXBSQCWlkRENsw+a3KRvn4\niWpp8ONrSGEt1sPeZ/uM9zLXuIbANGdjaSoWKWZpZZJDCeEEj54iIr/YybJjWduE\n62LHA5btAgMBAAECggEAH+VGHwY2tOJhj5eM1aOfWyQ79slPyxPWINpx9vczANfE\ncwb1xu0XY7TtnLzVRrBI44kUxNSVRK8Rpu7vRC6tLhbZEwPvr/Q+0lDeb1bpyNO1\nqqeexoDvCKIAadqyhOpUGl+j5ItiCGr0CicfNLdZEiv6cXgPAt0XbQhhfMLzl/cU\nsyhCU1Hd8GIIQDnnK3CTSiiJlJU2hljMHVIJWSEX3DtJm/j3mZQfYk93IFlUI1MA\nOSTi6THENZ/1MdlxH2DOk+PfFi383nGl6P7sWWKrUO+SSuV3Z8bhjHXwzd/CoB6T\nA8oVVosj+mbp02B0Ap9yDlEoIvGzJq4Mm8wvgDeuYQKBgQDcuxGumP5fiYPzlpbC\nTVTgy2O8FK/7ki0qnTzCcnmmpakXYtMA6dwjQnpW8OVdeKEygJMRL1z1rPTDqw3x\nsdglUz0dgNOEkq5/P0N456unc12k3zrFM9IzdmYzh5d5ysR8Xl3QMMFQwQcFQ+5A\nNFGbwqUeU9QtzT022gpBg5ncSQKBgQDcVb8b1JgErrK46QTVRXU+2gc5vx5ES6Rx\nbXl4UllHot26PjYQckPVO4/2fFTJKu0wr6edw9+SKQ4IYO2ocS0eNDIuBFku3GEe\n1G3QXs/itFqeSFVCCkXmiTA/4aujB5qEAs740G60ENx2j5vjASkriWls14kGSaFw\n50KtzmD9hQKBgDpawmv0SpubUWUelLC9nQjo/G8G0RejJ4mylBOcDAlAlpl2KO5+\n5RH1Sz6c5SZ287bUQw0yBlN07CimmkMhj1Ee1nNsUX8lADjn0sCuDrVwTHuAAJuN\n/a5ZSN+qoyMxtgxjLk4R9amRvndn5B7ZNhIFvX1tEBUjw2Ey968mSZDxAoGBANAw\nJRdC1TD3cN/PLUXnD1WH5XPm5c5aOtMCQdgy1zEc7qzfw23eycFdOjYIXISIDv4F\nuzcSsNkF+cBo9aZG6f60CwX4Ddx9VzcuOWS9cWggSc9tQUHZOxsNXY2+ydKNiK5b\niP0I1NFHbUiJgR4JJsGAYSD6tvo98FEh8psPeg2RAoGBAMbVzKqUyuSgiKYVlD/X\nwaQr1BBH8LppJUGExz+WiXc1WdY+Scm1DNRn6jEK/RfiGkzWZywRYURaHVy0zgVf\n/43Dlj9SHdUvwh2Kn3h+mf3WI1fPHuXJy7kAeYzB1Cl+60nZUeNQddZdVaK0MahG\nwiMZx/gRvBJisxR/1aI6jzFo\n-----END PRIVATE KEY-----\n",
  "client_email": "firebase-adminsdk-75cq4@itara-89760.iam.gserviceaccount.com",
  "client_id": "115765186157333358066",
  "auth_uri": "https://accounts.google.com/o/oauth2/auth",
  "token_uri": "https://oauth2.googleapis.com/token",
  "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
  "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/firebase-adminsdk-75cq4%40itara-89760.iam.gserviceaccount.com"
}`))}

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
