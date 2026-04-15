package algoritmosbusca

import "testing"

func TestBuscaLinear(t *testing.T) {
	arr := []int{15, 3, 9, 8, 22, 10, 5}

	t.Run("Elemento no meio", func(t *testing.T) {
		idx := BuscaLinear(8, arr)
		if idx != 3 {
			t.Errorf("Esperado índice 3, obtido %d", idx)
		}
	})

	t.Run("Elemento no início", func(t *testing.T) {
		idx := BuscaLinear(15, arr)
		if idx != 0 {
			t.Errorf("Esperado índice 0, obtido %d", idx)
		}
	})

	t.Run("Elemento no fim", func(t *testing.T) {
		idx := BuscaLinear(5, arr)
		if idx != 6 {
			t.Errorf("Esperado índice 6, obtido %d", idx)
		}
	})

	t.Run("Elemento inexistente", func(t *testing.T) {
		idx := BuscaLinear(99, arr)
		if idx != -1 {
			t.Errorf("Esperado índice -1, obtido %d", idx)
		}
	})

	t.Run("Array vazio", func(t *testing.T) {
		idx := BuscaLinear(10, []int{})
		if idx != -1 {
			t.Errorf("Esperado índice -1, obtido %d", idx)
		}
	})

	t.Run("Elementos duplicados", func(t *testing.T) {
		arrDup := []int{5, 10, 10, 20}
		idx := BuscaLinear(10, arrDup)
		if idx != 1 {
			t.Errorf("Esperado índice 1 (primeira ocorrência), obtido %d", idx)
		}
	})
}
