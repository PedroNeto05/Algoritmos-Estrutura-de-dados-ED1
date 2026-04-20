package list

import (
	"testing"
)

// factory para a LinkedList
func newLinkedList(t *testing.T) *LinkedList {
	t.Helper()
	return &LinkedList{}
}

func TestLinkedListImplementation(t *testing.T) {
	t.Run("Add and Size", func(t *testing.T) {
		l := newLinkedList(t)

		l.Add(10)
		l.Add(20)

		if l.Size() != 2 {
			t.Errorf("Esperado tamanho 2, obtido %d", l.Size())
		}
	})

	t.Run("Add Multiple Elements", func(t *testing.T) {
		l := newLinkedList(t)

		for i := range 50 {
			l.Add(i)
		}

		if l.Size() != 50 {
			t.Errorf("Esperado tamanho 50, obtido %d", l.Size())
		}

		for i := range 50 {
			val, err := l.Get(i)
			if err != nil || val != i {
				t.Errorf("Valor incorreto na posição %d: %d", i, val)
			}
		}
	})

	t.Run("Get Valid and Invalid", func(t *testing.T) {
		l := newLinkedList(t)
		l.Add(100)

		val, err := l.Get(0)
		if err != nil || val != 100 {
			t.Errorf("Get falhou: val=%d err=%v", val, err)
		}

		if _, err := l.Get(1); err == nil {
			t.Error("Esperava erro para índice fora do range")
		}
	})

	t.Run("Set Valid", func(t *testing.T) {
		l := newLinkedList(t)
		l.Add(1)

		if err := l.Set(99, 0); err != nil {
			t.Fatalf("Erro no Set: %v", err)
		}

		val, _ := l.Get(0)
		if val != 99 {
			t.Errorf("Esperado 99, obtido %d", val)
		}
	})

	t.Run("Set Invalid", func(t *testing.T) {
		l := newLinkedList(t)
		l.Add(1)

		if err := l.Set(10, -1); err == nil {
			t.Error("Esperava erro para índice negativo")
		}

		if err := l.Set(10, 5); err == nil {
			t.Error("Esperava erro para índice fora do range")
		}
	})

	t.Run("AddOnIndex Beginning", func(t *testing.T) {
		l := newLinkedList(t)

		l.Add(2)
		l.Add(3)

		if err := l.AddOnIndex(1, 0); err != nil {
			t.Fatalf("Erro ao inserir no início: %v", err)
		}

		val, _ := l.Get(0)
		if val != 1 {
			t.Errorf("Esperado 1, obtido %d", val)
		}
	})

	t.Run("AddOnIndex Middle", func(t *testing.T) {
		l := newLinkedList(t)

		l.Add(1)
		l.Add(3)

		if err := l.AddOnIndex(2, 1); err != nil {
			t.Fatalf("Erro ao inserir no meio: %v", err)
		}

		val, _ := l.Get(1)
		if val != 2 {
			t.Errorf("Esperado 2, obtido %d", val)
		}
	})

	t.Run("AddOnIndex Invalid", func(t *testing.T) {
		l := newLinkedList(t)
		l.Add(1)

		if err := l.AddOnIndex(10, -1); err == nil {
			t.Error("Esperava erro para índice negativo")
		}

		if err := l.AddOnIndex(10, 5); err == nil {
			t.Error("Esperava erro para índice fora do range")
		}
	})

	t.Run("Delete Beginning", func(t *testing.T) {
		l := newLinkedList(t)

		l.Add(1)
		l.Add(2)

		if err := l.RemoveOnIndex(0); err != nil {
			t.Fatalf("Erro ao deletar início: %v", err)
		}

		val, _ := l.Get(0)
		if val != 2 {
			t.Errorf("Esperado 2, obtido %d", val)
		}
	})

	t.Run("Delete Middle", func(t *testing.T) {
		l := newLinkedList(t)

		l.Add(1)
		l.Add(2)
		l.Add(3)

		if err := l.RemoveOnIndex(1); err != nil {
			t.Fatalf("Erro ao deletar meio: %v", err)
		}

		val, _ := l.Get(1)
		if val != 3 {
			t.Errorf("Esperado 3, obtido %d", val)
		}
	})

	t.Run("Delete End", func(t *testing.T) {
		l := newLinkedList(t)

		l.Add(1)
		l.Add(2)
		l.Add(3)

		if err := l.RemoveOnIndex(2); err != nil {
			t.Fatalf("Erro ao deletar final: %v", err)
		}

		if l.Size() != 2 {
			t.Errorf("Esperado tamanho 2, obtido %d", l.Size())
		}
	})

	t.Run("Delete Invalid", func(t *testing.T) {
		l := newLinkedList(t)
		l.Add(1)

		if err := l.RemoveOnIndex(-1); err == nil {
			t.Error("Esperava erro para índice negativo")
		}

		if err := l.RemoveOnIndex(5); err == nil {
			t.Error("Esperava erro para índice fora do range")
		}
	})

	t.Run("Delete Empty List", func(t *testing.T) {
		l := newLinkedList(t)

		if err := l.RemoveOnIndex(0); err == nil {
			t.Error("Esperava erro ao deletar lista vazia")
		}
	})

	t.Run("Delete Until Empty", func(t *testing.T) {
		l := newLinkedList(t)

		l.Add(1)
		l.Add(2)

		if err := l.RemoveOnIndex(0); err != nil {
			t.Fatalf("Erro: %v", err)
		}

		if err := l.RemoveOnIndex(0); err != nil {
			t.Fatalf("Erro: %v", err)
		}

		if l.Size() != 0 {
			t.Errorf("Lista deveria estar vazia")
		}
	})

	t.Run("Get Empty List", func(t *testing.T) {
		l := newLinkedList(t)

		if _, err := l.Get(0); err == nil {
			t.Error("Esperava erro ao acessar lista vazia")
		}
	})

	t.Run("Stress Mixed Operations", func(t *testing.T) {
		l := newLinkedList(t)

		for i := range 100 {
			l.Add(i)
		}

		for range 50 {
			if err := l.RemoveOnIndex(0); err != nil {
				t.Fatalf("Erro no delete: %v", err)
			}
		}

		if l.Size() != 50 {
			t.Errorf("Esperado tamanho 50, obtido %d", l.Size())
		}

		for i := 0; i < l.Size(); i++ {
			_, err := l.Get(i)
			if err != nil {
				t.Errorf("Erro ao acessar índice %d", i)
			}
		}
	})

	t.Run("Reverse Even Number of Elements", func(t *testing.T) {
		l := newLinkedList(t)
		l.Add(1)
		l.Add(2)
		l.Add(3)
		l.Add(4)

		l.Reverse()

		expected := []int{4, 3, 2, 1}
		for i, exp := range expected {
			val, _ := l.Get(i)
			if val != exp {
				t.Errorf("Índice %d: esperado %d, obtido %d", i, exp, val)
			}
		}
	})

	t.Run("Reverse Odd Number of Elements", func(t *testing.T) {
		l := newLinkedList(t)
		l.Add(10)
		l.Add(20)
		l.Add(30)

		l.Reverse()

		expected := []int{30, 20, 10}
		for i, exp := range expected {
			val, _ := l.Get(i)
			if val != exp {
				t.Errorf("Índice %d: esperado %d, obtido %d", i, exp, val)
			}
		}
	})

	t.Run("Reverse Single Element", func(t *testing.T) {
		l := newLinkedList(t)
		l.Add(42)

		l.Reverse()

		val, _ := l.Get(0)
		if val != 42 {
			t.Errorf("Esperado 42, obtido %d", val)
		}
	})

	t.Run("Reverse Empty List", func(t *testing.T) {
		l := newLinkedList(t)

		// Não deve causar panic
		l.Reverse()

		if l.Size() != 0 {
			t.Errorf("Tamanho deveria continuar 0, obtido %d", l.Size())
		}
	})
}
