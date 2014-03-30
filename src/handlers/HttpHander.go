package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const base = "src/html"

func HttpHandler(w http.ResponseWriter, r *http.Request) {

	var path string
	if r.URL.Path == "/" {
		path = base + "/EniacMonopoly.html"
	} else {
		path = base + r.URL.Path
	}

	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("Invalid web request: ", err)
		return
	}

	w.Write(bytes)
}
