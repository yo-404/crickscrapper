package Fixture

func SaveToJson(){
	var choice string
	var filename string
	fmt.Println("Would you like to save the following data as JSON")
	fmt.Println("Yes")
	fmt.Println("No")

	switch choice{

	case "y","Y","YES","yes":
		fmt.Println("Please enter the file name ")
		fmt.Scan(&filename)
		prefix:="./"
		file:= prefix+filename
		suffix := ".json"
		if !strings.HasSuffix(file,suffix)
		file1 := file + suffix
		err = os.WriteFile(file1, data, 0644)
		fmt.Printf("File SUccessfully saved in the S./ directory as %v",file)
		if err != nil {
		panic(err)
		}
	case "n","N","NO","no":
		Exit(1)
	

	default:
	fmt.Println("Please input a valid choice!!")
	
	}
}