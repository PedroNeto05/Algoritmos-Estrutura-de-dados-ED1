package algoritmos

import "testing"

func TestBuscaBinaria(t *testing.T) {
	arr := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}

	t.Run("Elemento no meio exato", func(t *testing.T) {
		idx := BuscaBinaria(arr, 12, 0, len(arr))
		if idx != 5 {
			t.Errorf("Esperado índice 5, obtido %d", idx)
		}
	})

	t.Run("Elemento na metade inferior (esquerda)", func(t *testing.T) {
		idx := BuscaBinaria(arr, 4, 0, len(arr))
		if idx != 1 {
			t.Errorf("Esperado índice 1, obtido %d", idx)
		}
	})

	t.Run("Elemento na metade superior (direita)", func(t *testing.T) {
		idx := BuscaBinaria(arr, 18, 0, len(arr))
		if idx != 8 {
			t.Errorf("Esperado índice 8, obtido %d", idx)
		}
	})

	t.Run("Elemento no limite inferior (início)", func(t *testing.T) {
		idx := BuscaBinaria(arr, 2, 0, len(arr))
		if idx != 0 {
			t.Errorf("Esperado índice 0, obtido %d", idx)
		}
	})

	t.Run("Elemento no limite superior (fim)", func(t *testing.T) {
		idx := BuscaBinaria(arr, 20, 0, len(arr))
		if idx != 9 {
			t.Errorf("Esperado índice 9, obtido %d", idx)
		}
	})

	t.Run("Elemento inexistente (menor que o mínimo)", func(t *testing.T) {
		idx := BuscaBinaria(arr, 1, 0, len(arr))
		if idx != -1 {
			t.Errorf("Esperado índice -1, obtido %d", idx)
		}
	})

	t.Run("Elemento inexistente (maior que o máximo)", func(t *testing.T) {
		idx := BuscaBinaria(arr, 99, 0, len(arr))
		if idx != -1 {
			t.Errorf("Esperado índice -1, obtido %d", idx)
		}
	})

	t.Run("Elemento inexistente (no meio dos valores)", func(t *testing.T) {
		idx := BuscaBinaria(arr, 7, 0, len(arr))
		if idx != -1 {
			t.Errorf("Esperado índice -1, obtido %d", idx)
		}
	})

	t.Run("Array com 1 elemento (encontrado)", func(t *testing.T) {
		arrUnico := []int{5}
		idx := BuscaBinaria(arrUnico, 5, 0, len(arrUnico))
		if idx != 0 {
			t.Errorf("Esperado índice 0, obtido %d", idx)
		}
	})

	t.Run("Array com 1 elemento (não encontrado)", func(t *testing.T) {
		arrUnico := []int{5}
		idx := BuscaBinaria(arrUnico, 10, 0, len(arrUnico))
		if idx != -1 {
			t.Errorf("Esperado índice -1, obtido %d", idx)
		}
	})

	t.Run("Array vazio", func(t *testing.T) {
		arrVazio := []int{}
		idx := BuscaBinaria(arrVazio, 10, 0, -1)
		if idx != -1 {
			t.Errorf("Esperado índice -1, obtido %d", idx)
		}
	})

	t.Run("Edge Case - Convergência de ponteiros (Array de 2 elementos)", func(t *testing.T) {
		arrCurto := []int{10, 20}

		idx10 := BuscaBinaria(arrCurto, 10, 0, len(arrCurto))
		if idx10 != 0 {
			t.Errorf("Esperado índice 0 para o valor 10, obtido %d", idx10)
		}

		idx20 := BuscaBinaria(arrCurto, 20, 0, len(arrCurto))
		if idx20 != 1 {
			t.Errorf("Esperado índice 1 para o valor 20, obtido %d", idx20)
		}
	})
}
