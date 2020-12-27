package main

import "github.com/gin-gonic/gin"

import "net/http"
import (
    "context"
    "time"
    "log"
    "fmt"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type InsertOneRequest struct {
    ID              primitive.ObjectID `json:"id" bson:"_id"`
    Hoge            string             `json:"hoge" bson:"hoge"`
    Fuga            string             `json:"fuga" bson:"fuga"`
}

func connectDb() (*mongo.Client){
    client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://mongodb"))
    if err != nil {
        log.Fatal(err)
    }
    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
    err = client.Connect(ctx)
    if err != nil {
            log.Fatal(err)
    }
    return client
}

func createTestCollection(client *mongo.Client) {
    collection := client.Database("local").Collection("test")
    request := InsertOneRequest{
        ID: primitive.NewObjectID(),
        Hoge: "hoge",
        Fuga: "fuga",
    }
    response, err := collection.InsertOne(context.Background(), request)
    if err != nil {
        log.Fatalln(err)
    }
    fmt.Println(response)
}

func main() {
    client := connectDb()
    createTestCollection(client)
    engine:= gin.Default()
    engine.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "hello world",
        })
    })
    engine.Run(":3000")
}