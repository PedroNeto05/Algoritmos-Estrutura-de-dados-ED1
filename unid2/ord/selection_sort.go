package ord

func SelectionSort(arr []int) {
	tamanho := len(arr)

	for i := 0; i < tamanho-1; i++ {

		menorIndex := i

		for j := i + 1; j < tamanho; j++ {
			if arr[j] < arr[menorIndex] {
				menorIndex = j
			}
		}

		if menorIndex != i {
			arr[i], arr[menorIndex] = arr[menorIndex], arr[i]
		}
	}
}
