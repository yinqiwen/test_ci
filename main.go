package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func startHTTPServer(listen string) {
	http.HandleFunc("/webhook", func(w http.ResponseWriter, r *http.Request) {
		data, err := ioutil.ReadAll(r.Body)
		if nil != err {
			log.Printf("Invalid notify:%v", err)
			return
		}
		log.Printf("Read notify %s", string(data))
	})
	err := http.ListenAndServe(listen, nil)
	if nil != err {
		log.Printf("Failed to start HTTP server with err:%v", err)
	}
}
func main() {
	startHTTPServer(os.Args[1])
}
