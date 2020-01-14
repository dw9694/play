package main

import (
	"encoding/json"
	"fmt"
	"github.com/dimonchik0036/vk-api"
	"log"
	"net/http"
	"net/url"
	"os"
)

var client, _ = vkapi.NewClientFromToken(os.Getenv("VK_TOKEN"))

func VideoOptions(r *http.Request) url.Values {
	q := r.URL.Query()
	v := url.Values{}

	v.Set("q", q.Get("query"))
	v.Set("filters", q.Get("filters"))
	v.Set("longer", q.Get("longer"))
	v.Set("count", q.Get("count"))
	return v
}

func GetVideo(r *http.Request) []byte {
	v := VideoOptions(r)
	res, err := client.Do(vkapi.NewRequest("video.search", "", v))
	if err != nil {
		panic(err)
	}

	items, _ := json.Marshal(res)
	return items
}

func handler(w http.ResponseWriter, r *http.Request) {
	items := GetVideo(r)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	_, err := w.Write(items)
	if err != nil {
		panic(err)
	}
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Printf("[INFO] start server on port %d\n", 7531)
	log.Fatal(http.ListenAndServe(":7531", nil))
}
