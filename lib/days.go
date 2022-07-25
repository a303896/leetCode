package lib

import "fmt"

type Day int

const (
	SUN = iota
	MON
	TUES
	WED
	THUR
	FRI
	SAT
)

var week = [7]string{"Sunday", "Monday", "Tuesday", "Wednesday ", "Thursday", "Friday", "Saturday"}

func (d Day) String() string {
	return fmt.Sprintf("%d is %s", d, week[d])
}
