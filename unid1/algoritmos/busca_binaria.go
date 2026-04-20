package algoritmos

func BuscaBinariaRecursiva(arr []int, val, ini, fim int) int {
	meio := (fim + ini) / 2

	if ini <= fim {
		if arr[meio] == val {
			return meio
		} else if arr[meio] > val {
			return BuscaBinariaRecursiva(arr, val, ini, meio-1)
		} else {
			return BuscaBinariaRecursiva(arr, val, meio+1, fim)
		}
	}

	return -1
}
