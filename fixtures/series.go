package Fixture

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/gocolly/colly/v2"
	model "github.com/yo-404/crickscrapper/Model"
)

type Series struct {
	Name string `json:"series,omitempty"`
	URL  string `json:"url,omitempty"`
}

type SeriesInfo struct {
	Series string `json:"series"`
	URL    string `json:"url"`
}

func SearchBySeries() string {
	CrawlUrl := "https://www.espncricinfo.com/ci/content/match/fixtures_futures.html"
	c := colly.NewCollector()
	tours := []Series{}

	c.OnHTML(".divright > ul", func(h *colly.HTMLElement) {

		var index int = 0
		for i := 0; i < 7; i++ {
			stat := "li:nth-child(" + strconv.Itoa(index) + ")"
			stat1 := "li:nth-child(" + strconv.Itoa(index) + ") a"

			future1 := h.ChildText(stat)
			future := h.ChildAttr(stat1, "href")

			tour := Series{
				Name: future1,
				URL:  future,
			}
			tours = append(tours, tour)

			index++
		}

	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	err := c.Visit(CrawlUrl)
	if err != nil {
		log.Fatal(err)
	}

	data1, err := json.MarshalIndent(tours, " ", "")
	if err != nil {
		log.Fatal(err)
	}

	var seriesData []SeriesInfo
	err = json.Unmarshal([]byte(data1), &seriesData)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		// return
	}

	var filteredData []SeriesInfo
	for _, series := range seriesData {
		if series.Series != "" || series.URL != "" {
			filteredData = append(filteredData, series)
		}
	}

	filteredJSON, err := json.MarshalIndent(filteredData, " ", "")
	if err != nil {
		fmt.Println("Error serializing JSON:", err)
		// return
	}

	// fmt.Println(string(filteredJSON))

	err = os.WriteFile("./series.json", filteredJSON, 0644)
	if err != nil {
		panic(err)
	}

	// fmt.Println(filteredData[2].Series)
	// fmt.Println(filteredData[2].URL)

	// fmt.Println(seriesData)

	for i, info := range filteredData {
		fmt.Printf("%v Series: %s \n", i+1, info.Series)
	}

	var choice int
	fmt.Println("Enter your choice (1/2/3..) - ")
	fmt.Scan(&choice)
	seriesname := filteredData[choice-1].URL
	fmt.Println("Selected url is :", seriesname)
	return seriesname

}

func ScrapforSeries(seriesname string) string {
	baseURL := "https://www.espncricinfo.com"
	varURL := "/match-schedule-fixtures-and-results"
	mainURl := baseURL + seriesname
	if !strings.HasSuffix(mainURl, varURL) {
		mainURl = mainURl + varURL
	}
	fmt.Println(mainURl)
	return mainURl

}

func CrawlerForSeries(CrawlUrl string) {
	matches := []model.SeriesInfo{}
	c := colly.NewCollector()
	c.OnHTML("div[class=ds-flex]", func(h *colly.HTMLElement) {
		team1 := (h.ChildText(".ds-my-1:nth-child(1) p"))
		teamx := (h.ChildText("div.ds-flex.ds-items-center.ds-min-w-0.ds-mr-1>p"))
		// teamx := h.ChildAttr("div.ds-flex.ds-items-center.ds-min-w-0.ds-mr-1", "title")
		team2 := strings.ReplaceAll(teamx, team1, "")
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
		score := h.ChildText("div.ds-text-compact-s.ds-text-typo.ds-text-right.ds-whitespace-nowrap")
		score1 := strings.ReplaceAll(score, "&", "  ")

		i := model.SeriesInfo{
			Date:       h.ChildText("div.ds-text-compact-xs.ds-font-bold.ds-w-24"),
			Type:       h.ChildText("span.ds-text-tight-s.ds-font-medium.ds-text-typo"),
			Tournament: h.ChildText("span.ds-text-tight-s.ds-font-regular.ds-text-typo.ds-underline.ds-decoration-ui-stroke"),
			Location:   location,
			Result:     h.ChildText("p.ds-text-tight-s.ds-font-regular.ds-line-clamp-2.ds-text-typo"),
			Score:      score1,
			Time:       result,
			Team1:      team1,
			Team2:      team2,
		}

		matches = append(matches, i)

	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	err := c.Visit(CrawlUrl)
	if err != nil {
		log.Fatal(err)
	}

	data, err := json.MarshalIndent(matches, " ", "")
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile("./seriesinfo.json", data, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))
}
