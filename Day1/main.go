package main

import (
	"SoftwareGoDay1/data"
	"SoftwareGoDay1/humanity"
)

func main() {
	data.ReadFile("./test.csv")
	data.LineToCSV("abc,def,ghi")
	humanity.NewHumanFromCsvFile("./test.csv")
	return
}
