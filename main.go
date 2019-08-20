package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("https://yuedu.baidu.com/ebook/c4e9eeab6aec0975f46527d3240c844769eaa0bb?fr=aladdin&key=%E6%B8%B8%E6%B3%B3")
	if err != nil {
		panic(err)
	}

	if resp.StatusCode != 200 {
		fmt.Println("err")
	}

	dom, err := goquery.NewDocumentFromReader(resp.Body)


	if err != nil{
		log.Fatalln(err)
	}

	dom.Find("p").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Text())
	})
}