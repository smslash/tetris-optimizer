package valid

func IsO(slice [][]string) bool {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if slice[i][j] == "#" {
				if i+1 <= 3 && j+1 <= 3 {
					if slice[i][j+1] == "#" && slice[i+1][j] == "#" && slice[i+1][j+1] == "#" {
						return true
					}
				}
			}
		}
	}
	return false
}
