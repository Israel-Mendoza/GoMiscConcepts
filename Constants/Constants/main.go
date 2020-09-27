package main

import "fmt"

type Weekday int

const (
	_ Weekday = iota
	Sunday
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func (w Weekday) workDay() string {
	switch w {
	case 2, 3, 4, 5, 6:
		return "Go to work!"
	default:
		return "You can relax!"
	}
}

func (w Weekday) String() string {
	return [...]string{
		"Sunday",
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturday",
	}[w-1]
}

func main() {
	fmt.Printf("Today is %s. %s", Sunday.String(), Sunday.workDay())
}
