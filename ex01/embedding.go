package ex01

import "fmt"

type Human struct {
	fieldOne int16
	fieldSec string
}

func (h Human) someParentFunc() string {
	return fmt.Sprintf("%d %s", h.fieldOne, h.fieldSec)
}

type Action struct {
	Human
}
