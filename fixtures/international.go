package Fixture

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/gocolly/colly/v2"
	model "github.com/yo-404/crickscrapper/Model"
)

func ByTeamInternational(CrawlUrl string) ([]byte, []model.Matches) {

	matches := []model.Matches{}
	c := colly.NewCollector()
	c.OnHTML("div[class=ds-flex]", func(h *colly.HTMLElement) {
		team1 := (h.ChildText(".ds-my-1:nth-child(1) p"))
		teamx := (h.ChildText("div.ds-flex.ds-items-center.ds-min-w-0.ds-mr-1>p"))
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

		i := model.Matches{
			Date:       h.ChildText("div.ds-text-compact-xs.ds-font-bold.ds-w-24"),
			Type:       h.ChildText("span.ds-text-tight-s.ds-font-medium.ds-text-typo"),
			Tournament: h.ChildText("span.ds-text-tight-s.ds-font-regular.ds-text-typo.ds-underline.ds-decoration-ui-stroke"),
			Location:   location,
			Result:     h.ChildText("p.ds-text-tight-s.ds-font-regular.ds-line-clamp-2.ds-text-typo"),
			Score:      h.ChildText("div.ds-text-compact-s.ds-text-typo.ds-text-right.ds-whitespace-nowrap"),
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

	fmt.Println(string(data))
	return data, matches

}
