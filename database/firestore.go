package database

import (
	"context"
	"log"
	"os"

	firestore "cloud.google.com/go/firestore"
	"firebase.google.com/go"
	"google.golang.org/api/option"
)

func ConnectFirestore() *firestore.Client {
    credentialsPath := os.Getenv("FIREBASE_CREDENTIALS_PATH")

    ctx := context.Background()

    opt := option.WithCredentialsFile(credentialsPath)

    app, err := firebase.NewApp(ctx, nil, opt)
    if err != nil {
        log.Fatalf("error initializing Firestore app: %v", err)
    }

    client, err := app.Firestore(ctx)
    if err != nil {
        log.Fatalf("error getting Firestore client: %v", err)
    }

    return client
}
