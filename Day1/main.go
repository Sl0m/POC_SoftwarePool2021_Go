package main

import (
	"SoftwareGoDay1/humanity"
	"fmt"
)

func main() {
	//	data.ReadFile("./test.csv")
	//	data.LineToCSV("abc,def,ghi")
	//	humanity.NewHumanFromCsvFile("./test.csv")
	//	humanity.NewHumanFromJsonFile("./medium.json")
	fmt.Println(&humanity.Pilot{&humanity.Human{"Jason", 10, "Fr", true}})
	humanity.Preparer.Prepare(&humanity.Pilot{&humanity.Human{"Jason", 10, "Fr", false}})
	humanity.Preparer.Prepare(&humanity.Pilot{&humanity.Human{"Jason", 10, "Fr", true}})
}
