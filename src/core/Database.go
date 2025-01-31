package core

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


var MongoClient *mongo.Client


func ConnectDB() {
	
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017") 
	
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("No se pudo conectar a MongoDB:", err)
	}

	fmt.Println("Conexión exitosa a MongoDB")
	MongoClient = client
}


func GetCollection(collectionName string) *mongo.Collection {
	if MongoClient == nil {
		log.Fatal("Error: La conexión a MongoDB no ha sido inicializada. Llama a core.ConnectDB() antes de usar GetCollection().")
	}
	return MongoClient.Database("tu_base_de_datos").Collection(collectionName)
}
