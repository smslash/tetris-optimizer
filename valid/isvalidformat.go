package valid

func IsValidFormat(arr string) ([][]string, bool) {
	slice := make([][]string, 4)
	for i := range slice {
		slice[i] = make([]string, 4)
	}

	if len(arr) != 16 {
		return slice, false
	}

	count_hash := 0
	n := 0
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			slice[i][j] = string(arr[n])
			n++

			if slice[i][j] == "#" {
				count_hash++
			} else if slice[i][j] != "." {
				return slice, false
			}
		}
	}

	if count_hash != 4 {
		return slice, false
	}

	return slice, true
}
