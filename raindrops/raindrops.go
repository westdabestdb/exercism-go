package raindrops

import (
	"fmt"
)

//Convert function takes an integer as parameter, finds this number's factors, and based on factors, returns a string or returns the number based on conditions.
func Convert(number int) string {
	var drop string                //define an empty string
	for i := 1; i <= number; i++ { //run a loop from 1 to number,
		if number%i == 0 { //if index divides number without any reminder, check for the number
			if i == 3 { // if number is equal to 3
				drop += "Pling" //append this string to drop string
			} else if i == 5 { // else if number is equal to 5
				drop += "Plang" //append this string to drop string
			} else if i == 7 { // else if number is equal to 7
				drop += "Plong" //append this string to drop string
			} else { //if none of these matches the current index
				if i == number && drop == "" { //checks for if index equals to number and if there was no 3,5,7 as factors,
					drop = fmt.Sprint(i) //sets the drop string equal to the string value of number
				}
			}
		}
	}
	return drop

}
