package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {

	// 1.打开JSON 文件
	// 2. 开一个 http server
	// 3. 设置路由接收 story 参数
	// 4. 输出网页。并且链接到正确的 store 页面

	type AdventureOption struct {
		Text string `json:"text"`
		Arc  string `json:"arc"`
	}

	type Adventure struct {
		Title   string            `json:"title"`
		Story   []string          `json:"story"`
		Options []AdventureOption `json:"options"`
	}

	var chapters map[string]Adventure

	jsonFile, err := os.Open("story.json")

	if err != nil {
		log.Fatal("Open json file error")
	}

	contentByte, err := ioutil.ReadAll(jsonFile)

	err = json.Unmarshal(contentByte, &chapters)

	if err != nil {
		log.Fatal("JSON unmarshal failed")
	}

	fmt.Println(chapters["debate"])

	http.HandleFunc("/story", func(w http.ResponseWriter, r *http.Request) {
		// 这里需要渲染网站 html 模板
		type htmlVars struct {
			Adventure Adventure
			Arc       string
			Story     map[string]Adventure
		}

		arc := r.URL.Query().Get("adventure")

		if arc == "" {
			arc = "home"
		}

		adventure := chapters[arc]

		tmpl, err := template.ParseFiles("adventure.html")
		if err != nil {
			log.Fatal("Get adventure fail")
		}

		output := htmlVars{
			Adventure: adventure,
			Arc:       arc,
			Story:     chapters,
		}

		tmpl.Execute(w, output)
	})

	http.ListenAndServe(":3000", nil)

}
