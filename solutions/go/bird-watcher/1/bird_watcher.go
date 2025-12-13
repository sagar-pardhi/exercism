package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	// panic("Please implement the TotalBirdCount() function")
	birdCount := 0
	for i := 0; i < len(birdsPerDay); i++ {
		birdCount += birdsPerDay[i]
	}

	return birdCount
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	// panic("Please implement the BirdsInWeek() function")
	weeks := [][]int{}

	for i := 0; i < week; i++ {
		if len(birdsPerDay) > 0 {
			week := birdsPerDay[:7]
			birdsPerDay = birdsPerDay[len(week):]
			weeks = append(weeks, week)
		}
	}

	return TotalBirdCount(weeks[week-1])
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	// panic("Please implement the FixBirdCountLog() function")
	for i := 0; i < len(birdsPerDay); i += 2 {
		birdsPerDay[i] += 1
	}

	return birdsPerDay
}
