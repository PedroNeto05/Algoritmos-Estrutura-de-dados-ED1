package ord

func BubbleSort(arr []int) {
	tamanho := len(arr)

	for i := 0; i < tamanho-1; i++ {

		troca := false
		for j := 0; j < tamanho-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				troca = true
			}
		}

		if !troca {
			break
		}
	}
}
