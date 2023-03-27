package main

func lemonadeChange(bills []int) bool {
	five, ten, twenty := 0, 0, 0
	for _, val := range bills {
		if val == 5 {
			five++
		}
		if val == 10 {
			if five <= 0 {
				return false
			}
			five--
			ten++
		}
		if val == 20 {
			if ten > 0 && five > 0 {
				ten--
				five--
				twenty++

			} else if ten <= 0 && five >= 3 {
				twenty++
				five -= 3
			} else {
				return false
			}

		}
	}
	return true
}
