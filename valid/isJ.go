package valid

func IsJ(slice [][]string) bool {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if slice[i][j] == "#" {
				if i+2 <= 3 && j-1 >= 0 {
					if slice[i+1][j] == "#" && slice[i+2][j] == "#" && slice[i+2][j-1] == "#" {
						return true
					}
				}

				if j+2 <= 3 && i+1 <= 3 {
					if slice[i][j+1] == "#" && slice[i][j+2] == "#" && slice[i+1][j+2] == "#" {
						return true
					}
				}

				if i+2 <= 3 && j+1 <= 3 {
					if slice[i+1][j] == "#" && slice[i+2][j] == "#" && slice[i][j+1] == "#" {
						return true
					}
				}

				if i+1 <= 3 && j+2 <= 3 {
					if slice[i+1][j] == "#" && slice[i+1][j+1] == "#" && slice[i+1][j+2] == "#" {
						return true
					}
				}
			}
		}
	}
	return false
}
