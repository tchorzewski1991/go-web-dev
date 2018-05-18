package main

import (
	"net/http"
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

type urls struct {
	Path string  `yaml:"path"`
	Url  string  `yaml:"url"`
}

func main() {
	http.ListenAndServe(":8080", http.HandlerFunc(yamlHandler))
}

func yamlHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=`UTF-8`")

	if req.Method == http.MethodPost {
		http.Redirect(w, req, "/", http.StatusMovedPermanently)
	}

	if req.URL.Path == "/" {
		http.NotFound(w, req)
		return
	}

	f, err := ioutil.ReadFile("urls.yaml")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	mapping := []urls{}

	err = yaml.Unmarshal(f, &mapping)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, req, func (mapping []urls) string {
		mapper := make(map[string]string)

		for _, url := range mapping { mapper[url.Path] = url.Url }

		if destination, ok := mapper[req.URL.Path]; ok {
			return destination
		}

		return "/" }(mapping), http.StatusMovedPermanently)
}