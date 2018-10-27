package leap

// IsLeapYear takes integer type year variable, calculates if it's a leap year and returns a type boolean.
func IsLeapYear(year int) bool {
	return year%400 == 0 || (year%4 == 0 && year%100 != 0)
}
