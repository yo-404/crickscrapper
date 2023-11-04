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
	fmt.Println("4. Search by link ")
	fmt.Print("Enter your choice (1/2/3..): ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		seriesname := Fixture.SearchBySeries()
		CrawlUrl := Fixture.ScrapforSeries(seriesname)
		data, matches := Fixture.CrawlerForSeries(CrawlUrl)
		Fixture.SaveToFile(matches, data)

	case 2:
		CrawlURL := crawler.ScrapforTeam()
		data, matches := Fixture.ByTeamInternational(CrawlURL)
		Fixture.SaveToFile(matches, data)
	case 3:
		CrawlURL := "https://www.espncricinfo.com/series/icc-cricket-world-cup-2023-24-1367856/match-schedule-fixtures-and-results"
		data, matches := Fixture.CrawlerForSeries(CrawlURL)
		Fixture.SaveToFile(matches, data)
	case 4:
		fmt.Println("Please enter url \n eg-https://www.espncricinfo.com/series/icc-cricket-world-cup-2023-24-1367856 ")
		var url string
		fmt.Scan(&url)
		CrawlURL:=Fixture.HasSuffix(url,"/match-schedule-fixtures-and-results")
		data, matches := Fixture.CrawlerForSeries(CrawlURL)
		Fixture.SaveToFile(matches, data)
	default:
		fmt.Println("More Options comming soon ..Till then pls work out with the above options")
	}
}

func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s Scrapping took %v", name, time.Since(start))
	}
}
