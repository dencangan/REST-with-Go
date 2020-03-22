package main

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"context"
	"time"
	"fmt"
	)

func connectMongo(username string, password string) {
	
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://" + username + ":" + password + 
	"@personaldatabase-rgdhs.azure.mongodb.net/test?retryWrites=true&w=majority",))

	if err != nil { 
		log.Fatal(err) 
	} 

	err = client.Ping(context.TODO(), nil)

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Connected to MongoDB!")
}

func main() {
	connectMongo("exampleUsername", "examplePassword")

}
