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

func PrepareMissionPart(objs ...Preparer) error {
	for i := range objs {
		Preparer.Prepare(objs[i])
	}
	return nil
}

type Checker interface {
	Check() bool
}

func (h *Human) Check() bool {
	return h.Ready
}

func CheckMissionPart(objs ...Checker) bool {
	for i := range objs {
		Checker.Check(objs[i])
	}
	return true
}
