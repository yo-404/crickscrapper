package crawler

import "fmt"

func ScrapforTeam() string {
	baseURL := "https://www.espncricinfo.com/team/"
	varURL := "/match-schedule-fixtures-and-results"
	var Team string
	var choice int
	var mainURl string
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
	case 2:
		Team = "australia-2"
	case 3:
		Team = "england-1"
	case 4:
		Team = "afghanistan-40"
	case 5:
		Team = "ireland-29"
	case 6:
		Team = "bangladesh-25"
	case 7:
		Team = "new-zealand-5"
	case 8:
		Team = "pakistan-7"
	case 9:
		Team = "south-africa-3"
	case 10:
		Team = "sri-lanka-8"
	case 11:
		Team = "west-indies-4"
	case 12:
		Team = "zimbabwe-9"
	case 13:
		Team = "namibia-28"
	case 14:
		Team = "nepal-33"
	case 15:
		Team = "netherlands-15"
	case 16:
		Team = "oman-37"
	case 17:
		Team = "papua-new-guinea-20"
	case 18:
		Team = "scotland-30"
	case 19:
		Team = "united-arab-emirates-27"
	case 20:
		Team = "united-states-of-america-11"

	default:
		fmt.Println("Please select a valid choice!!")

	}
	mainURl = baseURL + Team + varURL
	fmt.Println(mainURl)
	return mainURl
}
