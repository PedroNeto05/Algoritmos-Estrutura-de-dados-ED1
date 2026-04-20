package algoritmos

import "testing"

func TestBuscaBinariaAsc(t *testing.T) {
	arr := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}

	t.Run("Elemento no meio exato", func(t *testing.T) {
		idx := BuscaBinariaRecursivaAsc(arr, 12, 0, len(arr)-1)
		if idx != 5 {
			t.Errorf("Esperado índice 5, obtido %d", idx)
		}
	})

	t.Run("Elemento na metade inferior (esquerda)", func(t *testing.T) {
		idx := BuscaBinariaRecursivaAsc(arr, 4, 0, len(arr)-1)
		if idx != 1 {
			t.Errorf("Esperado índice 1, obtido %d", idx)
		}
	})

	t.Run("Elemento na metade superior (direita)", func(t *testing.T) {
		idx := BuscaBinariaRecursivaAsc(arr, 18, 0, len(arr)-1)
		if idx != 8 {
			t.Errorf("Esperado índice 8, obtido %d", idx)
		}
	})

	t.Run("Elemento no limite inferior (início)", func(t *testing.T) {
		idx := BuscaBinariaRecursivaAsc(arr, 2, 0, len(arr)-1)
		if idx != 0 {
			t.Errorf("Esperado índice 0, obtido %d", idx)
		}
	})

	t.Run("Elemento no limite superior (fim)", func(t *testing.T) {
		idx := BuscaBinariaRecursivaAsc(arr, 20, 0, len(arr)-1)
		if idx != 9 {
			t.Errorf("Esperado índice 9, obtido %d", idx)
		}
	})

	t.Run("Elemento inexistente (menor que o mínimo)", func(t *testing.T) {
		idx := BuscaBinariaRecursivaAsc(arr, 1, 0, len(arr)-1)
		if idx != -1 {
			t.Errorf("Esperado índice -1, obtido %d", idx)
		}
	})

	t.Run("Elemento inexistente (maior que o máximo)", func(t *testing.T) {
		idx := BuscaBinariaRecursivaAsc(arr, 99, 0, len(arr)-1)
		if idx != -1 {
			t.Errorf("Esperado índice -1, obtido %d", idx)
		}
	})

	t.Run("Elemento inexistente (no meio dos valores)", func(t *testing.T) {
		idx := BuscaBinariaRecursivaAsc(arr, 7, 0, len(arr)-1)
		if idx != -1 {
			t.Errorf("Esperado índice -1, obtido %d", idx)
		}
	})

	t.Run("Array com 1 elemento (encontrado)", func(t *testing.T) {
		arrUnico := []int{5}
		idx := BuscaBinariaRecursivaAsc(arrUnico, 5, 0, len(arrUnico)-1)
		if idx != 0 {
			t.Errorf("Esperado índice 0, obtido %d", idx)
		}
	})

	t.Run("Array com 1 elemento (não encontrado)", func(t *testing.T) {
		arrUnico := []int{5}
		idx := BuscaBinariaRecursivaAsc(arrUnico, 10, 0, len(arrUnico)-1)
		if idx != -1 {
			t.Errorf("Esperado índice -1, obtido %d", idx)
		}
	})

	t.Run("Array vazio", func(t *testing.T) {
		arrVazio := []int{}
		idx := BuscaBinariaRecursivaAsc(arrVazio, 10, 0, -1)
		if idx != -1 {
			t.Errorf("Esperado índice -1, obtido %d", idx)
		}
	})

	t.Run("Edge Case - Convergência de ponteiros (Array de 2 elementos)", func(t *testing.T) {
		arrCurto := []int{10, 20}

		idx10 := BuscaBinariaRecursivaAsc(arrCurto, 10, 0, len(arrCurto)-1)
		if idx10 != 0 {
			t.Errorf("Esperado índice 0 para o valor 10, obtido %d", idx10)
		}

		idx20 := BuscaBinariaRecursivaAsc(arrCurto, 20, 0, len(arrCurto)-1)
		if idx20 != 1 {
			t.Errorf("Esperado índice 1 para o valor 20, obtido %d", idx20)
		}
	})
}

func TestBuscaBinariaDesc(t *testing.T) {
	// Array invertido: do maior para o menor
	arr := []int{20, 18, 16, 14, 12, 10, 8, 6, 4, 2}

	t.Run("Desc - Elemento no meio exato", func(t *testing.T) {
		// len(arr) é 10. Meio é (0+9)/2 = 4. arr[4] é 12.
		idx := BuscaBinariaRecursivaDesc(arr, 12, 0, len(arr)-1)
		if idx != 4 {
			t.Errorf("Esperado índice 4, obtido %d", idx)
		}
	})

	t.Run("Desc - Elemento na metade inferior (esquerda - valores maiores)", func(t *testing.T) {
		idx := BuscaBinariaRecursivaDesc(arr, 18, 0, len(arr)-1)
		if idx != 1 {
			t.Errorf("Esperado índice 1, obtido %d", idx)
		}
	})

	t.Run("Desc - Elemento na metade superior (direita - valores menores)", func(t *testing.T) {
		idx := BuscaBinariaRecursivaDesc(arr, 4, 0, len(arr)-1)
		if idx != 8 {
			t.Errorf("Esperado índice 8, obtido %d", idx)
		}
	})

	t.Run("Desc - Elemento no limite inferior (início - maior valor)", func(t *testing.T) {
		idx := BuscaBinariaRecursivaDesc(arr, 20, 0, len(arr)-1)
		if idx != 0 {
			t.Errorf("Esperado índice 0, obtido %d", idx)
		}
	})

	t.Run("Desc - Elemento no limite superior (fim - menor valor)", func(t *testing.T) {
		idx := BuscaBinariaRecursivaDesc(arr, 2, 0, len(arr)-1)
		if idx != 9 {
			t.Errorf("Esperado índice 9, obtido %d", idx)
		}
	})

	t.Run("Desc - Elemento inexistente (maior que o máximo)", func(t *testing.T) {
		idx := BuscaBinariaRecursivaDesc(arr, 99, 0, len(arr)-1)
		if idx != -1 {
			t.Errorf("Esperado índice -1, obtido %d", idx)
		}
	})

	t.Run("Desc - Elemento inexistente (menor que o mínimo)", func(t *testing.T) {
		idx := BuscaBinariaRecursivaDesc(arr, 1, 0, len(arr)-1)
		if idx != -1 {
			t.Errorf("Esperado índice -1, obtido %d", idx)
		}
	})

	t.Run("Desc - Elemento inexistente (no meio dos valores)", func(t *testing.T) {
		idx := BuscaBinariaRecursivaDesc(arr, 7, 0, len(arr)-1)
		if idx != -1 {
			t.Errorf("Esperado índice -1, obtido %d", idx)
		}
	})

	t.Run("Desc - Array com 1 elemento (encontrado)", func(t *testing.T) {
		arrUnico := []int{5}
		idx := BuscaBinariaRecursivaDesc(arrUnico, 5, 0, len(arrUnico)-1)
		if idx != 0 {
			t.Errorf("Esperado índice 0, obtido %d", idx)
		}
	})

	t.Run("Desc - Array com 1 elemento (não encontrado)", func(t *testing.T) {
		arrUnico := []int{5}
		idx := BuscaBinariaRecursivaDesc(arrUnico, 10, 0, len(arrUnico)-1)
		if idx != -1 {
			t.Errorf("Esperado índice -1, obtido %d", idx)
		}
	})

	t.Run("Desc - Array vazio", func(t *testing.T) {
		arrVazio := []int{}
		idx := BuscaBinariaRecursivaDesc(arrVazio, 10, 0, -1)
		if idx != -1 {
			t.Errorf("Esperado índice -1, obtido %d", idx)
		}
	})

	t.Run("Desc - Edge Case - Convergência de ponteiros (Array de 2 elementos)", func(t *testing.T) {
		arrCurto := []int{20, 10}

		idx20 := BuscaBinariaRecursivaDesc(arrCurto, 20, 0, len(arrCurto)-1)
		if idx20 != 0 {
			t.Errorf("Esperado índice 0 para o valor 20, obtido %d", idx20)
		}

		idx10 := BuscaBinariaRecursivaDesc(arrCurto, 10, 0, len(arrCurto)-1)
		if idx10 != 1 {
			t.Errorf("Esperado índice 1 para o valor 10, obtido %d", idx10)
		}
	})
}
