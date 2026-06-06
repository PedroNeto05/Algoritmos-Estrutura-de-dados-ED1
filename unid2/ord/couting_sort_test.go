package ord

import (
	"reflect"
	"testing"
)

func TestCountingSort(t *testing.T) {
	// A nossa tabela clássica, mas SEM o teste de números negativos
	testes := []struct {
		nome     string
		entrada  []int
		esperado []int
	}{
		{
			nome:     "Lista desordenada comum",
			entrada:  []int{64, 25, 12, 22, 11},
			esperado: []int{11, 12, 22, 25, 64},
		},
		{
			nome:     "Lista já ordenada",
			entrada:  []int{1, 2, 3, 4, 5},
			esperado: []int{1, 2, 3, 4, 5},
		},
		{
			nome:     "Lista em ordem reversa",
			entrada:  []int{9, 8, 7, 6, 5},
			esperado: []int{5, 6, 7, 8, 9},
		},
		{
			nome:     "Lista com elementos duplicados",
			entrada:  []int{3, 1, 4, 1, 5, 9, 2, 6, 5},
			esperado: []int{1, 1, 2, 3, 4, 5, 5, 6, 9},
		},
		{
			nome:     "Lista vazia",
			entrada:  []int{},
			esperado: []int{},
		},
		{
			nome:     "Lista com apenas um elemento",
			entrada:  []int{42},
			esperado: []int{42},
		},
	}

	for _, tt := range testes {
		t.Run(tt.nome, func(t *testing.T) {
			// Cópia para não alterar o original durante o teste
			entradaCopia := make([]int, len(tt.entrada))
			copy(entradaCopia, tt.entrada)

			// Executamos o Counting Sort (ele é in-place, modifica a cópia)
			CountingSort(entradaCopia)

			// Validamos o resultado
			if !reflect.DeepEqual(entradaCopia, tt.esperado) {
				t.Errorf("Falha no teste '%s': esperado %v, recebido %v", tt.nome, tt.esperado, entradaCopia)
			}
		})
	}
}
