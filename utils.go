package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"fmt"

	"github.com/meilisearch/meilisearch-go"
)

var client = meilisearch.NewClient(meilisearch.Config{
	Host: "http://127.0.0.1:7700",
})

// CreateIndex creates indexes
func CreateIndex(n string) bool {
	if _, err := client.Indexes().Create(meilisearch.CreateIndexRequest{UID: n}); err != nil {
		return false
	}

	return true
}

// AddMovies adds movies to our dataset
func AddMovies() {
	moviesJSON, _ := os.Open("movies.json")
	defer moviesJSON.Close()

	byteValue, _ := ioutil.ReadAll(moviesJSON)
	var movies []map[string]interface{}
	json.Unmarshal(byteValue, &movies)

	updateRes, _ := client.Documents("movies").AddOrUpdate(movies)
	fmt.Println(updateRes.UpdateID)
}
