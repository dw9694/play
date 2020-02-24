package main

import (
	"encoding/json"
	"github.com/dimonchik0036/vk-api"
	"log"
	"net/http"
	"net/url"
	"os"
)

var client, _ = vkapi.NewClientFromToken(os.Getenv("VK_TOKEN"))

func main() {
	http.HandleFunc("/video", handlerVideo)
	log.Printf("[INFO] start server on port %d\n", 7531)
	log.Fatal(http.ListenAndServe(":7531", nil))
}

func handlerVideo(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	v := url.Values{}
	v.Set("q", q.Get("q"))
	v.Set("filters", q.Get("filters"))
	v.Set("longer", q.Get("longer"))
	v.Set("count", q.Get("count"))

	res, err := client.Do(vkapi.NewRequest("video.search", "", v))
	if err != nil {
		panic(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	items, _ := json.Marshal(res)
	values, _ := json.Marshal(v)
	w.Write(items)
	log.Printf("[REQUEST] %s\n", values)
	log.Printf("[RESPONSE] %s\n", items)
}
