package algoritmos

func BuscaBinariaRecursivaAsc(arr []int, val, ini, fim int) int {
	meio := (fim + ini) / 2

	if ini <= fim {
		if arr[meio] == val {
			return meio
		} else if arr[meio] > val {
			return BuscaBinariaRecursivaAsc(arr, val, ini, meio-1)
		} else {
			return BuscaBinariaRecursivaAsc(arr, val, meio+1, fim)
		}
	}

	return -1
}

func BuscaBinariaRecursivaDesc(arr []int, val, ini, fim int) int {
	meio := (fim + ini) / 2

	if ini <= fim {
		if arr[meio] == val {
			return meio
		} else if arr[meio] < val {
			return BuscaBinariaRecursivaDesc(arr, val, ini, meio-1)
		} else {
			return BuscaBinariaRecursivaDesc(arr, val, meio+1, fim)
		}
	}

	return -1
}

func BuscaBinariaIterativaAsc(arr []int, val int) int {
	ini := 0
	fim := len(arr) - 1

	for ini <= fim {
		meio := (fim + ini) / 2

		if arr[meio] == val {
			return meio
		}

		if arr[meio] > val {
			fim = meio - 1
		} else {
			ini = meio + 1
		}
	}

	return -1
}

func BuscaBinariaIterativaDesc(arr []int, val int) int {
	ini := 0
	fim := len(arr) - 1

	for ini <= fim {
		meio := (fim + ini) / 2

		if arr[meio] == val {
			return meio
		}

		if arr[meio] < val {
			fim = meio - 1
		} else {
			ini = meio + 1
		}
	}

	return -1
}
