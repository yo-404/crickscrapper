package crawler

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

func ScrapforTeam() string {
	baseURL := "https://www.espncricinfo.com/team/"
	varURL := "/match-schedule-fixtures-and-results"
	var Team string
	var choice int
	var mainURl string
	// var Teamname string
	fmt.Println("Select Team for internation Team fixtures")
	fmt.Println("1. INDIA")
	fmt.Println("2. AUSTRALIA")
	fmt.Println("3. ENGLAND")
	fmt.Println("4. AFGHANISTAN")
	fmt.Println("5. BANGLADESH")
	fmt.Println("6. IRELAND")
	fmt.Println("7. NEW-ZEALAND")
	fmt.Println("8. PAKISTAN")
	fmt.Println("9. SOUTH AFRICA")
	fmt.Println("10. SRI LANKA")
	fmt.Println("11. WEST INDIES")
	fmt.Println("12. ZIMBABWE")
	fmt.Println("13. NAMIBIA")
	fmt.Println("14. NEPAL")
	fmt.Println("15. NETHERLANDS")
	fmt.Println("16. OMAN")
	fmt.Println("17. PAPUA NEW GUINEA")
	fmt.Println("18. SCOTLAND")
	fmt.Println("19. UAE")
	fmt.Println("20. USA")

	fmt.Print("Enter your choice (1/2/3..): ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		Team = "india-6"
		// Teamname = "India"
	case 2:
		Team = "australia-2"
		// Teamname = "Australia"
	case 3:
		Team = "england-1"
		// Teamname = "England"
	case 4:
		Team = "afghanistan-40"
		// Teamname = "Afghanistan"
	case 5:
		Team = "ireland-29"
		// Teamname = "Ireland"
	case 6:
		Team = "bangladesh-25"
		// Teamname = "Bangladesh"
	case 7:
		Team = "new-zealand-5"
		// Teamname = "New Zealand"
	case 8:
		Team = "pakistan-7"
		// Teamname = "Pakistan"
	case 9:
		Team = "south-africa-3"
		// Teamname = "South Africa"
	case 10:
		Team = "sri-lanka-8"
		// Teamname = "Sri Lanka"
	case 11:
		Team = "west-indies-4"
		// Teamname = "West Indies"
	case 12:
		Team = "zimbabwe-9"
		// Teamname = "Zimbabwe"
	case 13:
		Team = "namibia-28"
		// Teamname = "Namibia"
	case 14:
		Team = "nepal-33"
		// Teamname = "Nepal"
	case 15:
		Team = "netherlands-15"
		// Teamname = "Netherlands"
	case 16:
		Team = "oman-37"
		// Teamname = "Oman"
	case 17:
		Team = "papua-new-guinea-20"
		// Teamname = "Papua New Guinea"
	case 18:
		Team = "scotland-30"
		// Teamname = "Scotland"
	case 19:
		Team = "united-arab-emirates-27"
		// Teamname = "United Arab Emirates"
	case 20:
		Team = "united-states-of-america-11"
		// Teamname = "United States Of America"

	default:
		fmt.Println("Please select a valid choice!!")

	}
	mainURl = baseURL + Team + varURL
	fmt.Println(mainURl)
	return mainURl
}

func SelectSeriesType() string {
	FutureTournaments := "https://www.espncricinfo.com/ci/content/match/fixtures_futures.html"
	c := colly.NewCollector()
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	err := c.Visit(FutureTournaments)
	if err != nil {
		log.Fatal(err)
	}

	c.OnHTML("div.divleft", func(h *colly.HTMLElement) {
		fmt.Println(h.Text)
	})
	return FutureTournaments
}
