package valid

func IsTetromino(slice [][]string) bool {
	if IsJ(slice) || IsL(slice) || IsS(slice) || IsZ(slice) || IsT(slice) || IsI(slice) || IsO(slice) {
		return true
	}
	return false
}
