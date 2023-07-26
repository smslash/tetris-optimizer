package valid

func IsI(slice [][]string) bool {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if slice[i][j] == "#" {
				if j+3 <= 3 {
					if slice[i][j+1] == "#" && slice[i][j+2] == "#" && slice[i][j+3] == "#" {
						return true
					}
				}

				if i+3 <= 3 {
					if slice[i+1][j] == "#" && slice[i+2][j] == "#" && slice[i+3][j] == "#" {
						return true
					}
				}
			}
		}
	}
	return false
}
