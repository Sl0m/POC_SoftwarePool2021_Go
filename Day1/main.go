package main

import (
	"SoftwareGoDay1/humanity"
)

func main() {
//	data.ReadFile("./test.csv")
//	data.LineToCSV("abc,def,ghi")
//	humanity.NewHumanFromCsvFile("./test.csv")
	humanity.NewHumanFromJsonFile("./medium.json")
	return
}
