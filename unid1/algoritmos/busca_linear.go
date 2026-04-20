package algoritmos

func BuscaLinear(val int, arr []int) int {
	for i, v := range arr {
		if v == val {
			return i
		}
	}
	return -1
}
