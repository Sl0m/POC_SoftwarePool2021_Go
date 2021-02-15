package main

import (
	"SoftwareGoDay1/data"
	"SoftwareGoDay1/humanity"
)

func main() {
	data.ReadFile("./test.csv")
	data.LineToCSV("abc,def,ghi")
	testHuman := []string{"abc", "12", "ghi"}
	humanity.NewHumanFromCSV(testHuman)
	humanity.NewHumanFromCsvFile("./test.csv")
	return
}
