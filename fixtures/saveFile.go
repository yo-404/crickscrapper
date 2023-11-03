package Fixture

import (
	"fmt"

	model "github.com/yo-404/crickscrapper/Model"
)

var choice string
var ch int

func SaveToFile(xlsxdata []model.Matches, jsondata []byte) {
	fmt.Println("Would You like to save your results into file (y/n)")
	fmt.Scan(&choice)

	switch choice {
	case "y", "Y":
		fmt.Println("Select the format ")
		fmt.Println("1. json")
		fmt.Println("2. xlsx")
		fmt.Scan(&ch)

		switch ch {
		case 1:
			SaveToJson(jsondata)
		case 2:
			SaveToXLSX(xlsxdata)
		default:
			fmt.Println("Invalid Choice !! ")
		}
	}

}
