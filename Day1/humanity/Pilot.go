package humanity

import (
	"fmt"
	"strconv"
)

type Pilot struct {
	*Human
}

func (h *Human) String() string {
	return fmt.Sprintf("%v, %v years old from %v", h.Name, strconv.Itoa(h.Age), h.Country)
}
