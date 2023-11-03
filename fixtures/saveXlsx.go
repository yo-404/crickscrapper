package Fixture

import (
	"fmt"
	"strings"

	"github.com/tealeg/xlsx"
	model "github.com/yo-404/crickscrapper/Model"
)

func SaveToXLSX(data []model.Matches) {

	// var choice string
	var filename string
	fmt.Println("Please enter the file name ")
	fmt.Scan(&filename)
	prefix := "./"
	file := prefix + filename
	suffix := ".xlsx"
	filename = HasSuffix(file, suffix)
	file1 := strings.TrimPrefix(filename, "./")

	// Create a new XLS file
	xlsFile := xlsx.NewFile()
	sheet, err := xlsFile.AddSheet("Matches")
	if err != nil {
		fmt.Println("Error creating XLS sheet:", err)
		return
	}

	// Add headers
	row := sheet.AddRow()
	row.AddCell().Value = "Type"
	row.AddCell().Value = "Tournament"
	row.AddCell().Value = "Location"
	row.AddCell().Value = "Date"
	row.AddCell().Value = "Result"
	row.AddCell().Value = "Score"
	row.AddCell().Value = "Time"
	row.AddCell().Value = "Team1"
	row.AddCell().Value = "Team2"

	// Add data
	for _, match := range data {
		row = sheet.AddRow()
		row.AddCell().Value = match.Type
		row.AddCell().Value = match.Tournament
		row.AddCell().Value = match.Location
		row.AddCell().Value = match.Date
		row.AddCell().Value = match.Result
		row.AddCell().Value = match.Score
		row.AddCell().Value = match.Time
		row.AddCell().Value = match.Team1
		row.AddCell().Value = match.Team2
	}
	err = xlsFile.Save(filename)
	if err != nil {
		fmt.Println("Error saving XLS file:", err)
		return
	}

	fmt.Printf("XLS file saved as '%s'", file1)
}
