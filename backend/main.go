package main

import (
	"encoding/json"
	"github.com/dimonchik0036/vk-api"
	"log"
	"net/http"
	"os"
)

var client, _ = vkapi.NewClientFromToken(os.Getenv("VK_TOKEN"))

func main() {
	http.HandleFunc("/video", handlerVideo)

	log.Printf("[INFO] start server on port %d\n", 7531)
	log.Fatal(http.ListenAndServe(":7531", nil))
}

// Handle video.search
func handlerVideo(w http.ResponseWriter, r *http.Request) {
	res, err := client.Do(vkapi.NewRequest("video.search", "", r.URL.Query()))
	if err != nil {
		panic(err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	items, _ := json.Marshal(res)
	w.Write(items)
}
