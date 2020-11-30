package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"gopkg.in/yaml.v2"
)

type REDIRECTS []struct {
	Path string
	Url  string
}

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		// 1. 读取 yml 文件获取所有的跳转 map
		content, err := ioutil.ReadFile("redirects.yml")
		if err != nil {
			log.Fatal(err)
		}
		redirects := REDIRECTS{}
		err = yaml.Unmarshal([]byte(content), &redirects)
		if err != nil {
			log.Fatal(err)
		}
		// 2. 匹配参数中的 short url 得到 source url
		short := r.URL.Query().Get("short")
		for _, redirect := range redirects {
			if redirect.Path != short {
				continue
			}
			// 3. 进行 302 跳转
			http.Redirect(w, r, redirect.Url, 301)
		}

	})
	http.ListenAndServe(":3000", nil)
}
