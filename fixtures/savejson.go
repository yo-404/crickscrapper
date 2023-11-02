package Fixture

import (
	"fmt"
	"os"
	"strings"
)

func SaveToJson(data []byte) {
	var choice string
	var filename string
	fmt.Println("\n \n =================================================\n ")
	fmt.Println("Would you like to save the following data as JSON (y/n)")
	fmt.Scan(&choice)

	switch choice {
	case "y", "Y", "YES", "yes":
		fmt.Println("Please enter the file name ")
		fmt.Scan(&filename)
		prefix := "./"
		file := prefix + filename
		suffix := ".json"
		filename := HasSuffix(file, suffix)
		file1 := strings.TrimPrefix(filename, "./")
		err := os.WriteFile(filename, data, 0644)
		fmt.Printf("File Successfully saved in the ./ directory as %v \n \n ", file1)
		if err != nil {
			panic(err)
		}
	case "n", "N", "NO", "no":
		fmt.Println("OKAY..Closing now ...")
		return

	default:
		fmt.Println("Please input a valid choice!!")

	}
}

func HasSuffix(filename string, suffix string) string {
	if !strings.HasSuffix(filename, suffix) {
		filename = filename + suffix
	}
	return filename
}
