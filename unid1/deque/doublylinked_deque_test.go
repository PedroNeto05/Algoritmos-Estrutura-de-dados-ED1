package deque

import (
	"testing"
)

// Factory para facilitar a criação nos testes
func newLinkedDeque(t *testing.T) *LinkedDeque {
	t.Helper()
	return &LinkedDeque{}
}

func TestLinkedDeque(t *testing.T) {
	t.Run("Empty Bounds Constraints", func(t *testing.T) {
		d := newLinkedDeque(t)

		if _, err := d.DequeueFront(); err == nil {
			t.Error("Deveria retornar erro ao dar DequeueFront em deque vazio")
		}
		if _, err := d.DequeueRear(); err == nil {
			t.Error("Deveria retornar erro ao dar DequeueRear em deque vazio")
		}
		if _, err := d.Front(); err == nil {
			t.Error("Deveria retornar erro ao dar Front em deque vazio")
		}
		if _, err := d.Rear(); err == nil {
			t.Error("Deveria retornar erro ao dar Rear em deque vazio")
		}
	})

	t.Run("Queue Behavior (EnqueueRear -> DequeueFront)", func(t *testing.T) {
		d := newLinkedDeque(t)

		d.EnqueueRear(10)
		d.EnqueueRear(20)
		d.EnqueueRear(30)

		if d.Size() != 3 {
			t.Fatalf("Tamanho esperado 3, obtido %d", d.Size())
		}

		expected := []int{10, 20, 30} // FIFO
		for _, exp := range expected {
			val, err := d.DequeueFront()
			if err != nil {
				t.Fatalf("Erro inesperado no DequeueFront: %v", err)
			}
			if val != exp {
				t.Errorf("FIFO falhou: esperado %d, obtido %d", exp, val)
			}
		}

		if d.Size() != 0 {
			t.Errorf("Tamanho não zerou após esvaziar a lista: obtido %d", d.Size())
		}
	})

	t.Run("Stack Behavior (EnqueueFront -> DequeueFront)", func(t *testing.T) {
		d := newLinkedDeque(t)

		d.EnqueueFront(10)
		d.EnqueueFront(20)
		d.EnqueueFront(30)

		expected := []int{30, 20, 10} // LIFO
		for _, exp := range expected {
			val, err := d.DequeueFront()
			if err != nil {
				t.Fatalf("Erro inesperado no DequeueFront: %v", err)
			}
			if val != exp {
				t.Errorf("LIFO falhou: esperado %d, obtido %d", exp, val)
			}
		}
	})

	t.Run("Reverse Queue Behavior (EnqueueFront -> DequeueRear)", func(t *testing.T) {
		d := newLinkedDeque(t)

		d.EnqueueFront(1) // Fica no fundo (rear)
		d.EnqueueFront(2) // Fica no meio
		d.EnqueueFront(3) // Fica na frente (front)

		// DequeueRear tira do fundo
		expected := []int{1, 2, 3}
		for _, exp := range expected {
			val, err := d.DequeueRear()
			if err != nil {
				t.Fatalf("Erro inesperado no DequeueRear: %v", err)
			}
			if val != exp {
				t.Errorf("Esperado %d, obtido %d", exp, val)
			}
		}
	})

	t.Run("Single Element Edge Case (Pointers validation)", func(t *testing.T) {
		d := newLinkedDeque(t)

		d.EnqueueFront(99)

		f, err := d.Front()
		if err != nil {
			t.Fatalf("Erro inesperado no Front com 1 elemento: %v", err)
		}

		r, err := d.Rear()
		if err != nil {
			t.Fatalf("Erro inesperado no Rear com 1 elemento: %v", err)
		}

		if f != 99 || r != 99 {
			t.Errorf("Com 1 elemento, Front (%d) e Rear (%d) devem ser iguais a 99", f, r)
		}

		// Remove o único elemento validando o erro
		val, err := d.DequeueRear()
		if err != nil {
			t.Fatalf("Erro inesperado ao remover o último elemento com DequeueRear: %v", err)
		}
		if val != 99 {
			t.Errorf("Valor removido incorreto: esperado 99, obtido %d", val)
		}

		if !d.IsEmpty() {
			t.Error("Deque deveria estar vazio após a remoção")
		}

		// Garante que os ponteiros não ficaram apontando pro nada (Dangling pointers)
		if d.front != nil || d.rear != nil {
			t.Error("Ponteiros front e rear devem ser nil quando a lista esvazia")
		}
	})
}
