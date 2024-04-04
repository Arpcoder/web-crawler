package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gocolly/colly"
	"github.com/kr/pretty"
)

type details struct {
	Header string `json:"header"`
	Desc   string `json:"desc"`
	Href   string `json:"href"`
}

func main() {
	c := colly.NewCollector()

	allitems := []details{}

	// Find and visit all links
	c.OnHTML("div.three div.row div.column", func(e *colly.HTMLElement) {
		items := details{
			Header: e.ChildText("h2.ui"),
			Desc:   e.ChildText("p"),
			Href:   e.ChildAttr("a", "href"),
		}
		allitems = append(allitems, items)
	})

	c.Visit("http://go-colly.org/")
	pretty.Print(allitems)

	// Marshal the struct array into JSON bytes
	jsonData, err := json.Marshal(allitems)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	// Write JSON data to a file
	err = os.WriteFile("products.json", jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing JSON to file:", err)
		return
	}

	fmt.Println("JSON data written to products.json successfully")

}
