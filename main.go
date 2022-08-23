package main

import (
	"fmt"
	"switchcase/switchcase"
)

func main() {
	dayNumber := 2

	dayOfWeek := switchcase.GetDayOfWeek(dayNumber)

	inlineDayOfWeek := switchcase.InlineGetDayOfWeek(dayNumber)

	fmt.Println(dayOfWeek)
	fmt.Println(inlineDayOfWeek)
}
