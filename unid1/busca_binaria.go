package algoritmosbusca

func BuscaBinaria(arr []int, val, ini, fim int) int {
	meio := (fim + ini) / 2

	if ini < fim {
		if arr[meio] == val {
			return meio
		} else if arr[meio] > val {
			return BuscaBinaria(arr, val, ini, meio-1)
		} else {
			return BuscaBinaria(arr, val, meio+1, fim)
		}
	}

	return -1
}
