package switchcase

import "fmt"

func GetDayOfWeek(number int) string {
	fmt.Printf("The given number is: %d\n", number)

	switch number {
	case 1:
		return "Monday"
	case 2:
		return "Tuesday"
	case 3:
		return "Wednesday"
	case 4:
		return "Thursday"
	case 5:
		return "Friday"
	case 6:
		return "Saturday"
	case 7:
		return "Sunday"
	default:
		return "Invalid number"
	}
}

func InlineGetDayOfWeek(number int) string {
	fmt.Printf("The given number is: %d\n", number)

	switch {
	case number == 1:
		return "Monday"
	case number == 2:
		return "Tuesday"
	case number == 3:
		return "Wednesday"
	case number == 4:
		return "Thursday"
	case number == 5:
		return "Friday"
	case number == 6:
		return "Saturday"
	case number == 7:
		return "Sunday"
	default:
		return "Invalid number"
	}
}
