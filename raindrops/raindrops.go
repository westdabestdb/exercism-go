package raindrops

import (
	"strconv"
)

//Convert function takes an integer as parameter, finds this number's factors, and based on factors, returns a string or returns the number based on conditions.
func Convert(number int) string {
	var lookFor = []int{3, 5, 7}                            //defines a slice for the numbers we are looking for
	var raindropSpeak = []string{"Pling", "Plang", "Plong"} //defines a slice for strings corresponding to these numbers
	var drop string                                         //defines empty string
	for _i, v := range lookFor {                            //running a loop inside lookFor slice
		if number%v == 0 { //if value divides number without any reminder
			drop += raindropSpeak[_i] //appends corresponding raindropSpeak string to drop string
		}
	}
	if drop == "" { //if there is no factor
		drop = strconv.Itoa(number) //appends string value of number to drop string
	}
	return drop //returns drop
}
