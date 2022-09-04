package database

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func CreateClient() *firestore.Client {
	ctx := context.Background()
	sa := option.WithCredentialsFile("database/firekey.json")
	config := &firebase.Config{ProjectID: "toys-and-coys"}
	app, err := firebase.NewApp(ctx, config, sa)
	if err != nil {
		fmt.Println("error 1")
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		fmt.Println("error 2")
		log.Fatalln(err)
	}
	return client
}
