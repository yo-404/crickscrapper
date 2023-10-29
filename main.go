package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/yo-404/crickscrapper/crawler"

	"github.com/gocolly/colly/v2"
)

type Matches struct {
	Type       string `json:"type"`
	Tournament string `json:"tournament"`
	// Location string `json:"location"`
	Date   string `json:"date"`
	Result string `json:"result"`
	Time   string `json:"time"`
}

func main() {
	defer timer("main")()
	c := colly.NewCollector()

	matches := []Matches{}

	CrawlURL := crawler.ScrapforTeam()

	c.OnHTML("div[class=ds-flex]", func(h *colly.HTMLElement) {
		i := Matches{
			Date:       h.ChildText("div.ds-text-compact-xs.ds-font-bold.ds-w-24"),
			Type:       h.ChildText("span.ds-text-tight-s.ds-font-medium.ds-text-typo"),
			Tournament: h.ChildText("span.ds-text-tight-s.ds-font-regular.ds-text-typo.ds-underline.ds-decoration-ui-stroke"),
			Result:     h.ChildText("p.ds-text-tight-s.ds-font-regular.ds-line-clamp-2.ds-text-typo"),
			Time:       h.ChildText("span.ds-text-tight-xs"),

			// fmt.Println(Time)
			// Match := h.ChildAttr("span[]", "title")
			// fmt.Println(Match)
		}

		matches = append(matches, i)

	})

	c.OnHTML("div[class=ds-flex]", func(h *colly.HTMLElement) {
		// text := h.ChildText("div.ds-flex.div[class=ds-text-compact-xs.ds-font-bold.ds-w-24]")
		// fmt.Println(h.Text)
		// text := h.ChildText("div[class=ds-text-compact-xs ds-font-bold ds-w-24]")
		// fmt.Println(text)
		// date := h.ChildText("div.ds-text-compact-xs.ds-font-bold.ds-w-24")
		// fmt.Println(date)
		// match := h.ChildText()
		// fmt.Println(match)
		// Test := h.ChildAttr("a.ds-inline-flex.ds-items-start.ds-leading-none.!ds-inline", "title")
		// fmt.Println(Test)
	})

	// c.OnHTML("div[class=ds-flex]", func(h *colly.HTMLElement) {
	// test := h.ChildText("div.ds-text-tight-s.ds-font-regular.ds-truncate.ds-text-typo-mid3#text")
	// fmt.Println(test)
	// test := h.ChildText("span.ds-text-tight-s.ds-font-regular.ds-text-typo.ds-underline.ds-decoration-ui-stroke")
	// fmt.Println(test)
	// 	test := h.ChildText("p.ds-text-tight-m.ds-font-bold.ds-capitalize.ds-truncate")
	// 	i := 1
	// 	fmt.Printf(" Team %d : %v \t", i, test)
	// 	i++
	// })

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	err := c.Visit(CrawlURL)
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
