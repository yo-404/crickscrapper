package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gocolly/colly/v2"
)

type Matches struct {
	Type     string `json:"type"`
	Match    string `json:"match"`
	Location string `json:"location"`
	Time     string `json:"time"`
}

func main() {
	defer timer("main")()
	c := colly.NewCollector()

	// c.OnHTML("div.cb-col-100.cb-col ", func(h *colly.HTMLElement) {
	// text := h.ChildText("a.cb-col-33.cb-col.cb-mtchs-dy.text-bold")
	// fmt.Println(text)
	// match := h.ChildText("div.cb-ovr-flo.cb-col-60.cb-col.cb-mtchs-dy-vnu.cb-adjst-lst.span[itemprop=name] ")
	// fmt.Println(match)
	// match := h.ChildText("span[itemprop=name]")
	// fmt.Println(match)

	// location := h.ChildText("div.cb-font-12.text-gray.cb-ovr-flo")
	// fmt.Println(location)
	// time := h.ChildText("div.cb-col-40.cb-col.cb-mtchs-dy-tm.cb-adjst-lst")
	// fmt.Println(time)

	// })

	matches := []Matches{}

	c.OnHTML("div.cb-col-100.cb-col", func(h *colly.HTMLElement) {
		i := Matches{
			Type:     h.ChildText("a.cb-col-33.cb-col.cb-mtchs-dy.text-bold"),
			Match:    h.ChildText("div.cb-ovr-flo.cb-col-60.cb-col.cb-mtchs-d-vnu.cb-adjst-lst"),
			Location: h.ChildText("span[itemprop=name]"),
			Time:     h.ChildText("div.cb-col-40.cb-col.cb-mtchs-dy-tm.cb-adjst-lst"),
		}
		matches = append(matches, i)
	})

	// c.OnHTML("div.cb-col-67.cb-col", func(h *colly.HTMLElement) {
	// 	// match := h.ChildAttr("a", "title")
	// 	// fmt.Println(match)
	// 	match := h.ChildText("div.cb-ovr-flo.cb-col-60.cb-col.cb-mtchs-dy-vnu.cb-adjst-lst")
	// 	fmt.Println(match)
	// 	// location := h.ChildText("div.cb-font-12.text-gray.cb-ovr-flo")
	// 	// fmt.Println(location)
	// })

	// Set error handler
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	err := c.Visit("https://www.cricbuzz.com/cricket-schedule/upcoming-series/international")
	if err != nil {
		log.Fatal(err)
	}

	data, err := json.MarshalIndent(matches, " ", "")
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile("./matches.json", data, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))

}

func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s Scrapping took %v", name, time.Since(start))
	}
}

// c.OnHTML("div.cb-col-100.cb-col ", func(h *colly.HTMLElement) {
// 	text := h.ChildText("a.cb-col-33.cb-col.cb-mtchs-dy.text-bold")
// 	fmt.Println(text)
// 	match := h.ChildText("div.cb-ovr-flo.cb-col-60.cb-col.cb-mtchs-dy-vnu.cb-adjst-lst")
// 	fmt.Println(match)

// 	location := h.ChildText("div.cb-font-12.text-gray.cb-ovr-flo")
// 	fmt.Println(location)
// 	time := h.ChildText("div.cb-col-40.cb-col.cb-mtchs-dy-tm.cb-adjst-lst")
// 	fmt.Println(time)

// })

// items = append(items,i)

// text := h.ChildText("a.cb-col-33.cb-col.cb-mtchs-dy.text-bold")
// fmt.Println(text)
// match := h.ChildText("div.cb-ovr-flo.cb-col-60.cb-col.cb-mtchs-dy-vnu.cb-adjst-lst")
// fmt.Println(match)

// location := h.ChildText("div.cb-font-12.text-gray.cb-ovr-flo")
// fmt.Println(location)
// time := h.ChildText("div.cb-col-40.cb-col.cb-mtchs-dy-tm.cb-adjst-lst")
// fmt.Println(time)

// c.OnHTML("div.cb-col-100.cb-col ", func(h *colly.HTMLElement) {
// 	i := Matches{
// 		Type: h.ChildText("a.cb-col-33.cb-col.cb-mtchs-dy.text-bold"),
// 		Match: h.ChildText("div.cb-ovr-flo.cb-col-60.cb-col.cb-mtchs-dy-vnu.cb-adjst-lst")
// 		Location: h.ChildText("div.cb-font-12.text-gray.cb-ovr-flo"),
// 		Time: h.ChildText("div.cb-col-40.cb-col.cb-mtchs-dy-tm.cb-adjst-lst"),
// 	}
// 	items = append(items, i)
// })
