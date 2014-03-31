package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

const base = "src/html"

var mimeTypes = make(map[string]string)

func loadMimeTypes() {
	mimeTypes[".css"] = "text/css"
	mimeTypes[".js"] = "application/javascript"
	mimeTypes[".html"] = "text/html"
}

func HttpHandler(w http.ResponseWriter, r *http.Request) {

	if len(mimeTypes) == 0 {
		loadMimeTypes()
	}

	fmt.Println("Incoming web request:", r.URL.Path)

	var path string
	if r.URL.Path == "/" {
		path = base + "/EniacMonopoly.html"
	} else {
		path = base + r.URL.Path
	}

	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("Invalid web request:", err)
		return
	}

	mime, ok := mimeTypes[filepath.Ext(path)]

	if ok {
		w.Header().Add("Content-Type", mime)
	}
	w.Write(bytes)
}
