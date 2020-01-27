package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const assets = "./web"

type httpHandler = func(w http.ResponseWriter, r *http.Request)

// const proxyURL = "http://chat:8081"

func handleGet(w http.ResponseWriter, r *http.Request, c *mongo.Collection) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cur, err := c.Find(ctx, bson.D{})
	defer cur.Close(ctx)

	if err != nil {
		fmt.Println("Error finding products collection", err)
		return
	}

	var results []bson.M
	for cur.Next(ctx) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal("Error fetching products from collectin", err)
			return
		}
		results = append(results, result)
	}

	fmt.Println(results)
	fmt.Println(len(results))

	json.NewEncoder(w).Encode(results)

}

type user struct {
	FirstName string `json:"firstname"`
	Surname   string `json:"surname"`
	Email     string `json:"email"`
}

type product struct {
	MongoID     string `json:"_id"`
	ID          string `json:"id"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
	Quantity    int    `json:"quantity"`
}

type order struct {
	user
	id          string
	dateCreated string
	Products    []product `json:"products"`
}

func handlePost(w http.ResponseWriter, r *http.Request, c *mongo.Collection) {

	var u order
	json.NewDecoder(r.Body).Decode(&u)
	// check stock in mong....
	//...
	u.dateCreated = time.Now().UTC().Format("02-01-2006")
	u.id = uuid.New().String()

}

func handler(c *mongo.Collection) httpHandler {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")

		switch r.Method {
		case "GET":
			handleGet(w, r, c)
		case "POST":
			handlePost(w, r, c)
		}
	}
}

func main() {

	mongoURL := "mongodb://" + os.Getenv("MONOG_HOST") + ":27017"

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))

	if err != nil {
		fmt.Println("Error with client connecting", err)
		return
	}

	ctx, _ = context.WithTimeout(context.Background(), 2*time.Second)
	err = client.Ping(ctx, readpref.Primary())

	if err != nil {
		fmt.Println("Error pinging", err)
		return
	}

	db := client.Database("shop")
	collection := db.Collection("products")

	http.HandleFunc("/", handler(collection))

	fmt.Println("LISTENING ON 8080 - 3")

	http.ListenAndServe(":8080", nil)

}
