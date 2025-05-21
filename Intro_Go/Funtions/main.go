package main

//  return value

func sub(x int, y int) int {
	return x - y
}

func concat(s1 string, s2 string) string {
	return s1 + s2
}

func add(x, y int) int {
	return x + y
}

//	multiple return values and ignore one value

func getPoint() (int, int) {
	return 3, 4
}

func getName(firstName string, lastName string) (string, string) {
	return firstName, lastName
} // first_name, _ := getName("John", "Wall")

// named return values

func getCoord() (x int, y int) {
	// x and y are initialized with zero values

	return // automatically return x and y
}

func yearsUntilEvents(age int) (yearsUntilAdult int, yearsUntilDrinking int, yearsUntilCarRental int) {
	yearsUntilAdult = 18 - age
	if yearsUntilAdult < 0 {
		yearsUntilAdult = 0
	}
	yearsUntilDrinking = 21 - age
	if yearsUntilDrinking < 0 {
		yearsUntilDrinking = 0
	}
	yearsUntilCarRental = 25 - age
	if yearsUntilCarRental < 0 {
		yearsUntilCarRental = 0
	}
	return
}

func main() {

}
