package crawler

import (
	"fmt"
	"net/url"
)

func ScrapforTeam(team string) {
	parsedURL, err := url.Parse("https://www.espncricinfo.com/team/teamname/match-schedule-fixtures-and-results")
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}

	queryValues := parsedURL.Query()
	queryValues.Set("teamname", team)

	parsedURL.RawQuery = queryValues.Encode()
	modifiedURL := parsedURL.String()
	fmt.Println("Modified URL:", modifiedURL)

}
