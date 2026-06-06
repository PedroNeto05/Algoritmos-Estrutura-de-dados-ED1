package ord

func CountingSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	maior := arr[0]

	for _, v := range arr {
		if v > maior {
			maior = v
		}
	}

	coutingArr := make([]int, maior+1)

	for _, v := range arr {
		coutingArr[v]++
	}

	indexOriginal := 0

	for num, qtd := range coutingArr {
		for qtd > 0 {
			arr[indexOriginal] = num
			indexOriginal++
			qtd--
		}
	}
}
