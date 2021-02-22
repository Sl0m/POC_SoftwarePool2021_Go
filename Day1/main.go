package main

import (
	"SoftwareGoDay1/humanity"
)

func main() {
	//	data.ReadFile("./test.csv")
	//	data.LineToCSV("abc,def,ghi")
	//	humanity.NewHumanFromCsvFile("./test.csv")
	//	humanity.NewHumanFromJsonFile("./medium.json")
	pilotList := []humanity.Preparer{
		&humanity.Pilot{
			Human: &humanity.Human{
				Name:    "Jason",
				Age:     25,
				Country: "USA",
				Ready:   false,
			},
		},
		&humanity.Pilot{
			Human: &humanity.Human{
				Name:    "Jean-René",
				Age:     77,
				Country: "France",
				Ready:   false,
			},
		},
	}
	humanity.PrepareMissionPart(pilotList[0])
	humanity.PrepareMissionPart(pilotList[0])
	humanity.PrepareMissionPart(pilotList[1])
	humanity.PrepareMissionPart(pilotList[1])
}
