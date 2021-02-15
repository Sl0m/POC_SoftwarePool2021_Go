package humanity

import "fmt"

type Preparer interface {
	Prepare() error
}

func (h *Human) Prepare() error {
	if h.Ready == true {
		fmt.Printf("%v is ready !\n", h)
	}
	h.Ready = true
	return nil
}
