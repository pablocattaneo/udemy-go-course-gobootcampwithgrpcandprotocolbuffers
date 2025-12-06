package main

func main() {
	day := 3

	switch day {
	case 1:
		println("Monday")
	case 2:
		println("Tuesday")
	case 3:
		println("Wednesday")
	case 4:
		println("Thursday")
	case 5:
		println("Friday")
	case 6:
		println("Saturday")
	case 7:
		println("Sunday")
	default:
		println("Invalid day")
	}

	// using operators in switch

	numberForSwitchOperatorExample := 15

	switch {
	case numberForSwitchOperatorExample < 10:
		println("Less than 10")
	case numberForSwitchOperatorExample >= 10 && numberForSwitchOperatorExample <= 20:
		println("Between 10 and 20")
	case numberForSwitchOperatorExample > 20:
		println("Greater than 20")
	}

	number := 2

	switch number {
	case 1, 3, 5, 7, 9:
		println("Odd number")
	case 2, 4, 6, 8, 10:
		println("Even number")
	default:
		println("Number is out of range")
	}

	// fallthrough example
	grade := 'B'

	switch grade {
	case 'A':
		println("Excellent!")
	case 'B':
		println("Well done")
		fallthrough
	case 'C':
		println("You passed")
	case 'D':
		println("Better try again")
	default:
		println("Invalid grade")
	}
}
