package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/yinqiwen/gotoolkit/xjson"
)

func startHTTPServer(listen string) {
	http.HandleFunc("/webhook", func(w http.ResponseWriter, r *http.Request) {
		data, err := ioutil.ReadAll(r.Body)
		if nil != err {
			log.Printf("Invalid notify:%v", err)
			return
		}
		log.Printf("Read notify %s", string(data))
		x, err := xjson.Decode(bytes.NewBuffer(data))
		if nil != err {
			log.Printf("Read notify json error %v", err)
			return
		}
		if x.RGet("commits").ArraySize() > 0 {
			repo := x.RGet("repository").RGet("name").GetString()
			cmd := exec.Command("sh", repo + "_tag.sh")
			log.Printf("Running %s_tag.sh and waiting for it to finish...", repo)
			err = cmd.Run()
			log.Printf("Command finished with error: %v", err)
		} else {
			log.Printf("Not a commit.")
		}
	})
	err := http.ListenAndServe(listen, nil)
	if nil != err {
		log.Printf("Failed to start HTTP server with err:%v", err)
	}
}
func main() {
	startHTTPServer(os.Args[1])
}
