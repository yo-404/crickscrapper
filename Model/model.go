package model

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

type SeriesInfo struct {
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
