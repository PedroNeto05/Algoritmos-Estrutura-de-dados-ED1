package ord2

func CountingSort(arr []int) {
	max := arr[0]

	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	coutingArr := make([]int, max+1)

	for i := 0; i < len(arr); i++ {
		coutingArr[arr[i]]++
	}

	index := 0

	for v, i := range coutingArr {
		for i > 0 {
			arr[index] = v
			index++
			i--
		}
	}
}
