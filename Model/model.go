package model

type Matches struct {
	Type       string `json:"match"`
	Tournament string `json:"tournament"`
	Location   string `json:"location"`
	Date       string `json:"date"`
	Result     string `json:"result"`
	Score      string `json:"score"`
	Time       string `json:"time"`
	Team1      string `json:"team1"`
	Team2      string `json:"team2"`
}

type Series struct {
	Name string `json:"series,omitempty"`
	URL  string `json:"url,omitempty"`
}

type SeriesInfo struct {
	Series string `json:"series"`
	URL    string `json:"url"`
}
