package Fixture

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func SaveToJson(data []byte) {

	if _, err := os.Stat("output"); os.IsNotExist(err) {
		err := os.MkdirAll("output", 0777)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Directory created: output \n")
	}

	// var choice string
	var filename string
	fmt.Println("Please enter the file name ")
	fmt.Scan(&filename)
	prefix := "./output/"
	file := prefix + filename
	suffix := ".json"
	filename = HasSuffix(file, suffix)
	file1 := strings.TrimPrefix(filename, "./output/")

	err := os.WriteFile(filename, data, 0644)
	fmt.Printf("File Successfully saved in the output directory as %v \n \n ", file1)
	if err != nil {
		panic(err)
	}
}

func HasSuffix(filename string, suffix string) string {
	if !strings.HasSuffix(filename, suffix) {
		filename = filename + suffix
	}
	return filename
}
