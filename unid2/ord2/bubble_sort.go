package ord2

func BubbleSort(arr []int) {
	tamanho := len(arr)
	for i := 0; i < tamanho-1; i++ {
		teveTroca := false
		for j := 0; j < len(arr)-1-i; i++ {
			if arr[j] > arr[j+1] {
				arr[j+1], arr[j] = arr[j], arr[j+1]
				teveTroca = true
			}
		}
		if !teveTroca {
			break
		}
	}
}
