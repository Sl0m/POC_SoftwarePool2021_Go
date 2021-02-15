package data

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func ReadFile(path string) ([]string, error) {
	dat, err := ioutil.ReadFile(path)
	str := strings.Split(string(dat), "\n")
	return str, err
}

func LineToCSV(line string) ([]string, error) {
	lines := strings.Split(line, ",")
	if len(lines) == 0 {
		return nil, fmt.Errorf("error parsing line")
	}
	return lines, nil
}
