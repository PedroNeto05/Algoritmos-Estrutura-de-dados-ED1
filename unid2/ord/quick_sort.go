package ord

import "math/rand"

func QuickSort(arr []int) {
	quickSortRecursivo(arr, 0, len(arr)-1)
}

func quickSortRecursivo(arr []int, inicio, fim int) {
	if inicio < fim {
		indicePivo := partition(arr, inicio, fim)

		quickSortRecursivo(arr, inicio, indicePivo-1)
		quickSortRecursivo(arr, indicePivo+1, fim)
	}
}

func partition(arr []int, inicio, fim int) int {
	indiceAle := rand.Intn(fim-inicio+1) + inicio

	arr[indiceAle], arr[fim] = arr[fim], arr[indiceAle]

	pivo := arr[fim]

	i := inicio - 1

	for j := inicio; j < fim; j++ {
		if arr[j] < pivo {
			i++
			arr[j], arr[i] = arr[i], arr[j]

		}
	}

	arr[i+1], arr[fim] = arr[fim], arr[i+1]

	return i + 1
}
