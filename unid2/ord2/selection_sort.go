package ord2

func SelectionSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		menorIndex := 0
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[menorIndex] {
				menorIndex = j
			}
		}
		if menorIndex != i {
			arr[i], arr[menorIndex] = arr[menorIndex], arr[i]
		}
	}
}
