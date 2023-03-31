package pets

func CalculateYears(years int) (result [3]int) {
	result[0] = years

	for i := 1; i <= years; i++ {
		switch i {
		case 1:
			result[1] += 15
			result[2] += 15
		case 2:
			result[1] += 9
			result[2] += 9
		default:
			result[1] += 4
			result[2] += 5
		}
	}

	return
}
