package main

import (
    "context"
    "time"
    "log"
    "fmt"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "github.com/gin-gonic/gin"
    "net/http"
    "github.com/gin-contrib/cors"
)

type InsertOneRequest struct {
    ID              primitive.ObjectID `json:"id" bson:"_id"`
    Hoge            string             `json:"hoge" bson:"hoge"`
    Fuga            string             `json:"fuga" bson:"fuga"`
}

func connectDb() (*mongo.Client){
    client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://mongodb"))
    // client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
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
    engine := gin.Default()

    // CORS対策
    engine.Use(cors.New(cors.Config{
        AllowOrigins: []string{
            "http://localhost:3001",
            "http://easy-shrimp.com",
        },
        AllowMethods: []string{
            "POST",
            "GET",
            "OPTIONS",
            "DELETE",
        },
        AllowHeaders: []string{
            "Access-Control-Allow-Credentials",
            "Access-Control-Allow-Headers",
            "Content-Type",
            "Content-Length",
            "Accept-Encoding",
            "Authorization",
        },
        AllowCredentials: false,
    }))
    engine.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "🦐🦐🦐",
        })
    })
    engine.Run(":3000")
}