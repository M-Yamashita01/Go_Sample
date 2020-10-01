package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Item struct {
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
}

func main() {
	response, err := http.Get("http://qiita.com/api/v2/users/M-Yamashii/items?page=1&per_page=10")
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var data []Item

	if err := json.Unmarshal(body, &data); err != nil {
		log.Fatal(err)
	}

	for _, item := range data {
		fmt.Printf("%s, %s\n", item.CreatedAt, item.Title)
	}
}
