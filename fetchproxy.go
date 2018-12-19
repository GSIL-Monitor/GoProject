package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func fetchProxy() {
	//doc, err := goquery.NewDocument("https://cn-proxy.com")
	doc, err := goquery.NewDocument("https://cn-proxy.com/archives/218")
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find("tbody tr").Each(func(i int, s *goquery.Selection) {
		var ip, port string
		s.Children().Each(func(i int, selection *goquery.Selection) {
			switch i {
			case 0:
				ip = selection.Text()
			case 1:
				port = selection.Text()
			default:
			}
		})
		fmt.Println(strings.Split(ip, "\n")[0] + ":" + strings.Split(port, "\n")[0])
	})
}

func main() {
	fetchProxy()
}
