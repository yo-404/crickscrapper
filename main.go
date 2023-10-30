package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/yo-404/crickscrapper/crawler"

	"github.com/gocolly/colly/v2"
)

type Matches struct {
	Type       string `json:"match"`
	Tournament string `json:"tournament"`
	Location   string `json:"location"`
	Date       string `json:"date"`
	Result     string `json:"result"`
	Score      string `json:"score"`
	Time       string `json:"time"`
	Opponent   string `json:"opponent"`
}

func main() {
	defer timer("main")()
	c := colly.NewCollector()

	matches := []Matches{}

	CrawlURL, Teamname := crawler.ScrapforTeam()

	c.OnHTML("div[class=ds-flex]", func(h *colly.HTMLElement) {
		team1 := (h.ChildText("div.ds-flex.ds-items-center.ds-min-w-0.ds-mr-1"))
		team2 := strings.ReplaceAll(team1, Teamname, "")
		result := h.ChildText("span.ds-text-tight-xs")
		if result == "RESULT" {
			result = "RESULT DECLARED"
		}

		// location
		var result1 string
		location := h.ChildText("div.ds-text-tight-s.ds-font-regular.ds-truncate.ds-text-typo-mid3")
		parts := strings.Split(location, "•")
		if len(parts) > 1 {
			result1 = strings.TrimSpace(parts[1])
		}
		for i := 0; i < len(result1); i++ {
			if result1[i] == ',' {
				location = result1[0:i]
				break
			}

		}

		i := Matches{
			Date:       h.ChildText("div.ds-text-compact-xs.ds-font-bold.ds-w-24"),
			Type:       h.ChildText("span.ds-text-tight-s.ds-font-medium.ds-text-typo"),
			Tournament: h.ChildText("span.ds-text-tight-s.ds-font-regular.ds-text-typo.ds-underline.ds-decoration-ui-stroke"),
			Location:   location,
			Result:     h.ChildText("p.ds-text-tight-s.ds-font-regular.ds-line-clamp-2.ds-text-typo"),
			Score:      h.ChildText("div.ds-text-compact-s.ds-text-typo.ds-text-right.ds-whitespace-nowrap"),
			Time:       result,
			Opponent:   team2,
		}

		matches = append(matches, i)

	})

	c.OnHTML("div[class=ds-flex]", func(h *colly.HTMLElement) {
		// team1 := h.ChildAttr("div.ds-flex.ds-items-center.ds-min-w-0.ds-mr-1", "title") != "India"
		// if h.ChildText("div.ds-flex.ds-items-center.ds-min-w-0.ds-mr-1") != "India" && h.ChildText("div.ds-flex.ds-items-center.ds-min-w-0.ds-mr-1") != "IND" {
		// 	team1 := (h.ChildText("div.ds-flex.ds-items-center.ds-min-w-0.ds-mr-1"))
		// 	team2 := strings.ReplaceAll(team1, "India", "")
		// 	fmt.Println(team2)
		// var result string
		// location := h.ChildText("div.ds-text-tight-s.ds-font-regular.ds-truncate.ds-text-typo-mid3")
		// parts := strings.Split(location, "•")
		// if len(parts) > 1 {
		// 	result = strings.TrimSpace(parts[1])
		// 	// fmt.Println(result)
		// }
		// for i := 0; i < len(result); i++ {
		// 	if result[i] == ',' {
		// 		fmt.Println(result[0:i])
		// 		break
		// 	}

		// }
		// score := h.ChildText("div.ds-text-compact-s.ds-text-typo.ds-text-right.ds-whitespace-nowrap")
		// fmt.Println(score)
	})

	// fmt.Println(Test)

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
