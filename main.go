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
	// Type     string `json:"type"`
	// Match    string `json:"match"`
	// Location string `json:"location"`
	Time string `json:"time"`
}

func main() {
	defer timer("main")()
	c := colly.NewCollector()

	matches := []Matches{}

	c.OnHTML("div[class=ds-flex]", func(h *colly.HTMLElement) {
		i := Matches{
			h.ChildText("div.ds-text-compact-xs.ds-font-bold.ds-w-24"),
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
		date := h.ChildText("div.ds-text-compact-xs.ds-font-bold.ds-w-24")
		fmt.Println(date)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	err := c.Visit("https://www.espncricinfo.com/team/india-6/match-schedule-fixtures-and-results")
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
