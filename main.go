package main

import (
	"fmt"
	"time"

	"github.com/yo-404/crickscrapper/crawler"
	Fixture "github.com/yo-404/crickscrapper/fixtures"
)

func main() {
	defer timer("main")()

	var choice int
	fmt.Println("Select How would you like to scrap the data")
	fmt.Println("1. By Series Name")
	fmt.Println("2. By International Teams")
	fmt.Println("3. World Cup 2023")
	fmt.Print("Enter your choice (1/2/3..): ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		seriesname := Fixture.SearchBySeries()
		CrawlUrl := Fixture.ScrapforSeries(seriesname)
		Fixture.CrawlerForSeries(CrawlUrl)
	case 2:
		CrawlURL, Teamname := crawler.ScrapforTeam()
		Fixture.ByTeamInternational(Teamname, CrawlURL)
	case 3:
		CrawlURL:="https://www.espncricinfo.com/series/icc-cricket-world-cup-2023-24-1367856/match-schedule-fixtures-and-results"
		Fixture.CrawlerForSeries(CrawlURL)

	default:
		fmt.Println("More Options comming soon ..Till then pls work out with the above options")
	}
	// seriesname := Fixture.SearchBySeries()
	// CrawlUrl := Fixture.ScrapforSeries(seriesname)
	// fmt.Println(CrawlUrl)
	// Fixture.CrawlerForSeries(CrawlUrl)
	// CrawlURL, Teamname := crawler.ScrapforTeam()
	// Fixture.ByTeamInternational(Teamname, CrawlURL)

}

func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s Scrapping took %v", name, time.Since(start))
	}
}
