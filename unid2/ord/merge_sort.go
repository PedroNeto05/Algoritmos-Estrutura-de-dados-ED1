package ord

func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	tamanho := len(arr)
	meio := tamanho / 2

	esquerda := MergeSort(arr[:meio])
	direita := MergeSort(arr[meio:])

	return merge(esquerda, direita)
}

func merge(esquerda, direita []int) []int {
	tamEsquerda := len(esquerda)
	tamDireita := len(direita)
	tamanhoTotal := tamEsquerda + tamDireita
	resultado := make([]int, tamanhoTotal)

	i, j, k := 0, 0, 0

	for i < len(esquerda) && j < len(direita) {
		if esquerda[i] < direita[j] {
			resultado[k] = esquerda[i]
			i++
		} else {
			resultado[k] = direita[j]
			j++
		}
		k++
	}

	for i < tamEsquerda {
		resultado[k] = esquerda[i]
		i++
		k++
	}

	for j < tamDireita {
		resultado[k] = direita[j]
		j++
		k++
	}

	return resultado
}
