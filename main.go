package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"gopkg.in/mgo.v2/bson"
)

//Define the port for the
const (
	WEBPORT = ":80"
)

//RespBody is Create the struct to handle the response json
type RespBody struct {
	Time struct {
		Updated    string    `json:"updated"`
		UpdatedISO time.Time `json:"updatedISO"`
		Updateduk  string    `json:"updateduk"`
	} `json:"time"`
	Disclaimer string `json:"disclaimer"`
	ChartName  string `json:"chartName"`
	Bpi        struct {
		USD struct {
			Code        string  `json:"code"`
			Symbol      string  `json:"symbol"`
			Rate        string  `json:"rate"`
			Description string  `json:"description"`
			RateFloat   float64 `json:"rate_float"`
		} `json:"USD"`
		GBP struct {
			Code        string  `json:"code"`
			Symbol      string  `json:"symbol"`
			Rate        string  `json:"rate"`
			Description string  `json:"description"`
			RateFloat   float64 `json:"rate_float"`
		} `json:"GBP"`
		EUR struct {
			Code        string  `json:"code"`
			Symbol      string  `json:"symbol"`
			Rate        string  `json:"rate"`
			Description string  `json:"description"`
			RateFloat   float64 `json:"rate_float"`
		} `json:"EUR"`
	} `json:"bpi"`
}

func main() {
	fmt.Println("My Application is starting")

	router := mux.NewRouter()

	http.Handle("/", router)

	router.HandleFunc("/home", homefunc)
	router.HandleFunc("/solconsumer", SolaceConsumer)

	http.ListenAndServe(WEBPORT, nil)

}

func homefunc(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "We have received request 2")

}

func SolaceConsumer(w http.ResponseWriter, r *http.Request) {

	reqbody, err := ioutil.ReadAll(r.Body)
	if err != nil {

		log.Print(err)
	}

	u := RespBody{}

	json.Unmarshal(reqbody, &u)

	MongoDBInsertFun(&u)

	fmt.Println(u)

}

func MongoDBInsertFun(body *RespBody) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://svcrm:svcrm@cluster0.rr5av.mongodb.net/techcody?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(databases)

	database := client.Database("techcody")
	podcastsCollection := database.Collection("training")

	insertResult, err := podcastsCollection.InsertOne(ctx, body)
	if err != nil {
		panic(err)
	}
	fmt.Println(insertResult.InsertedID)

}
