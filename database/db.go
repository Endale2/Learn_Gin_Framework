package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	
)
var DB *mongo.Database 


func  ConnectDatabase(){
	clientOptions:=options.Client().ApplyURI("mongodb://localhost:27017")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	client,err :=mongo.Connect(ctx, clientOptions)

	if err!=nil{
		log.Fatal("MongoDB   connection error..", err)
	}


	DB = client.Database("todo_app")

	log.Println("MongoDB  connected!")


}