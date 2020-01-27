package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path/filepath"
)

const assets = "./web"

func main() {

	path, err := filepath.Abs(assets)
	proxyURL := "http://" + os.Getenv("API_HOST")
	url, err := url.Parse(proxyURL)

	if err != nil {
		fmt.Println("cannot parse url", proxyURL)
		return
	}

	proxy := httputil.NewSingleHostReverseProxy(url)

	fmt.Println("URLLL:", url)

	if err != nil {
		fmt.Println("Damn cannot find path", assets, err)
	}

	fs := http.FileServer(http.Dir(path))

	http.Handle("/", fs)

	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {

		proxy.ServeHTTP(w, r)

	})

	http.ListenAndServe(":8086", nil)

}
