package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

const (
	WEBPORT = ":8080"
)

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

	fmt.Fprint(w, "We have received request")

}

func SolaceConsumer(w http.ResponseWriter, r *http.Request) {

	reqbody, err := ioutil.ReadAll(r.Body)
	if err != nil {

		log.Print(err)
	}

	u := RespBody{}

	json.Unmarshal(reqbody, &u)

	fmt.Println(u)

}
