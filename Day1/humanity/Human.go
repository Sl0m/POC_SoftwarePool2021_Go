package humanity

import (
	"SoftwareGoDay1/data"
	"fmt"
	"strconv"
)

//Human is the Astronaut Type
type Human struct {
	Name    string
	Age     int
	Country string
}

//NewHumanFromCSV create a Human from CSV
func NewHumanFromCSV(csv []string) (*Human, error) {
	var err error
	hum := Human{Name: csv[0], Country: csv[2]}
	hum.Age, err = strconv.Atoi(csv[1])
	fmt.Printf("%#v\n", hum)
	return &hum, err
}

func NewHumanFromCsvFile(path string) (*[]Human, error) {
	lines, err := data.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error while reading file")
	}
	var humans []Human
	for i := range lines {
		h, err := data.LineToCSV(lines[i])
		if err != nil {
			return nil, fmt.Errorf("error parsing line")
		}
		hum, _ := NewHumanFromCSV(h)
		humans = append(humans, *hum)
	}
	if err != nil {
		return nil, fmt.Errorf("error in human creation")
	}
	return &humans, nil
}