package Fixture

import (
	"fmt"
	"os"
	"strings"
)

func SaveToJson(data []byte) {
	// var choice string
	var filename string
	fmt.Println("Please enter the file name ")
	fmt.Scan(&filename)
	prefix := "./"
	file := prefix + filename
	suffix := ".json"
	filename = HasSuffix(file, suffix)
	file1 := strings.TrimPrefix(filename, "./")
	err := os.WriteFile(filename, data, 0644)
	fmt.Printf("File Successfully saved in the ./ directory as %v \n \n ", file1)
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
