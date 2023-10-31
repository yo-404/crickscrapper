package main

import (
	"fmt"
	"time"

	"github.com/yo-404/crickscrapper/crawler"
	international "github.com/yo-404/crickscrapper/fixtures"
)

func main() {
	defer timer("main")()
	CrawlURL, Teamname := crawler.ScrapforTeam()
	international.ByTeamInternational(Teamname, CrawlURL)

}

func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s Scrapping took %v", name, time.Since(start))
	}
}

// func SelectSeries() {
// 	d := colly.NewCollector()
// 	var tours []struct {
// 		Name string
// 		Link string
// 	}
// 	d.OnRequest(func(r *colly.Request) {
// 		fmt.Println("Visiting", r.URL)
// 	})

// 	d.OnHTML("div.divleft.li", func(e *colly.HTMLElement) {
// 		// var choice int

// 		name := e.ChildText("e.Text")
// 		link := e.ChildAttr("ul", "li")
// 		fmt.Println(name)
// 		fmt.Println(link)
// 		d.Visit(e.Request.AbsoluteURL(link))
// 		tour := struct {
// 			Name string
// 			Link string
// 		}{
// 			Name: name,
// 			Link: link,
// 		}
// 		tours = append(tours, tour)

// 		// fmt.Println("Enter the tournament whose fixtures you want to see!")
// 		// fmt.Scan(&choice)
// 		// fmt.Println(tour.Link[2])
// 	})

// 	d.OnError(func(r *colly.Response, err error) {
// 		log.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
// 	})

// 	err := d.Visit("https://www.espncricinfo.com/ci/content/match/fixtures_futures.html#second")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	for _, tour := range tours {
// 		fmt.Printf("Tour: %s\nLink: https://www.espncricinfo.com%s\n\n", tour.Name, tour.Link)
// 	}

// 	data1, err := json.MarshalIndent(tours, " ", "")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(string(data1))
// }
