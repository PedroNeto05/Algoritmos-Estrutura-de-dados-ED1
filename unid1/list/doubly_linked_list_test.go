package list

import (
	"testing"
)

func newDoublyLinkedList(t *testing.T) *DoublyLinkedList {
	t.Helper()
	return &DoublyLinkedList{}
}

func TestDoublyLinkedListImplementation(t *testing.T) {
	t.Run("Add and Size", func(t *testing.T) {
		l := newDoublyLinkedList(t)

		l.Add(10)
		l.Add(20)

		if l.Size() != 2 {
			t.Errorf("Esperado tamanho 2, obtido %d", l.Size())
		}
	})

	t.Run("Add Multiple Elements", func(t *testing.T) {
		l := newDoublyLinkedList(t)

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
		l := newDoublyLinkedList(t)
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
		l := newDoublyLinkedList(t)
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
		l := newDoublyLinkedList(t)
		l.Add(1)

		if err := l.Set(10, -1); err == nil {
			t.Error("Esperava erro para índice negativo")
		}

		if err := l.Set(10, 5); err == nil {
			t.Error("Esperava erro para índice fora do range")
		}
	})

	t.Run("AddOnIndex Beginning", func(t *testing.T) {
		l := newDoublyLinkedList(t)

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
		l := newDoublyLinkedList(t)

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

	t.Run("Delete Beginning", func(t *testing.T) {
		l := newDoublyLinkedList(t)

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
		l := newDoublyLinkedList(t)

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
		l := newDoublyLinkedList(t)

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

	t.Run("Stress Mixed Operations", func(t *testing.T) {
		l := newDoublyLinkedList(t)

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
	})

	t.Run("DLL: Reverse Traversal (Prev Pointers)", func(t *testing.T) {
		l := newDoublyLinkedList(t)
		l.Add(10)
		l.Add(20)
		l.Add(30)

		if l.tail == nil || l.tail.val != 30 {
			t.Fatalf("Tail incorreto. Esperado 30, obtido %v", l.tail)
		}

		curr := l.tail
		expectedValues := []int{30, 20, 10}

		for i, expected := range expectedValues {
			if curr == nil {
				t.Fatalf("Nó é nil prematuramente na iteração %d (reversa)", i)
			}
			if curr.val != expected {
				t.Errorf("Esperado %d, obtido %d na travessia reversa", expected, curr.val)
			}
			curr = curr.prev
		}

		if curr != nil {
			t.Errorf("Esperado nil antes do head, mas obteve nó com valor %d", curr.val)
		}
	})

	t.Run("DLL: Head and Tail Integrity", func(t *testing.T) {
		l := newDoublyLinkedList(t)
		l.Add(5)

		if l.head != l.tail {
			t.Error("Head e Tail devem apontar para o mesmo nó quando há 1 elemento")
		}
		if l.head.prev != nil {
			t.Error("Prev do Head deve ser nil")
		}
		if l.tail.next != nil {
			t.Error("Next do Tail deve ser nil")
		}

		l.Add(10)
		if l.head == l.tail {
			t.Error("Head e Tail devem apontar para nós diferentes agora")
		}
		if l.head.prev != nil {
			t.Error("Prev do Head deve continuar nil após inserção")
		}
		if l.tail.next != nil {
			t.Error("Next do Tail deve continuar nil após inserção")
		}
	})

	t.Run("DLL: Prev Pointer Integrity on AddOnIndex", func(t *testing.T) {
		l := newDoublyLinkedList(t)
		l.Add(10)
		l.Add(30)

		err := l.AddOnIndex(20, 1)
		if err != nil {
			t.Fatalf("Falha ao adicionar: %v", err)
		}

		node30 := l.tail
		if node30.val != 30 {
			t.Fatalf("Tail deveria ser 30, é %d", node30.val)
		}

		node20 := node30.prev
		if node20 == nil || node20.val != 20 {
			t.Fatalf("Prev do 30 deveria ser 20")
		}

		node10 := node20.prev
		if node10 == nil || node10.val != 10 {
			t.Fatalf("Prev do 20 deveria ser 10")
		}
	})

	t.Run("DLL: Prev Pointer Integrity on RemoveOnIndex", func(t *testing.T) {
		l := newDoublyLinkedList(t)
		l.Add(1)
		l.Add(2)
		l.Add(3)

		err := l.RemoveOnIndex(1)
		if err != nil {
			t.Fatalf("Falha ao remover: %v", err)
		}

		if l.head.val != 1 {
			t.Errorf("Head deveria ser 1, é %d", l.head.val)
		}
		if l.tail.val != 3 {
			t.Errorf("Tail deveria ser 3, é %d", l.tail.val)
		}

		if l.head.next != l.tail {
			t.Errorf("O next do head (1) deveria apontar para o tail (3)")
		}
		if l.tail.prev != l.head {
			t.Errorf("O prev do tail (3) deveria apontar para o head (1)")
		}
	})

	t.Run("DLL: Remove Tail Update", func(t *testing.T) {
		l := newDoublyLinkedList(t)
		l.Add(100)
		l.Add(200)

		err := l.RemoveOnIndex(1)
		if err != nil {
			t.Fatalf("Erro ao remover último elemento: %v", err)
		}

		if l.tail == nil || l.tail.val != 100 {
			t.Errorf("O tail não foi atualizado corretamente. Esperado 100")
		}
		if l.tail.next != nil {
			t.Errorf("O next do tail deve ser nil")
		}
	})
	t.Run("Reverse Even Number of Elements", func(t *testing.T) {
		l := newDoublyLinkedList(t)
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
		l := newDoublyLinkedList(t)
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
		l := newDoublyLinkedList(t)
		l.Add(42)

		l.Reverse()

		val, _ := l.Get(0)
		if val != 42 {
			t.Errorf("Esperado 42, obtido %d", val)
		}

		if l.head != l.tail || l.head.val != 42 {
			t.Errorf("Head e Tail deveriam apontar para o 42. Head: %v, Tail: %v", l.head, l.tail)
		}
	})

	t.Run("Reverse Empty List", func(t *testing.T) {
		l := newDoublyLinkedList(t)

		l.Reverse() // Não deve causar panic

		if l.Size() != 0 {
			t.Errorf("Tamanho deveria continuar 0, obtido %d", l.Size())
		}
		if l.head != nil || l.tail != nil {
			t.Errorf("Head e Tail deveriam continuar nil")
		}
	})

	t.Run("Reverse Head and Tail Integrity", func(t *testing.T) {
		l := newDoublyLinkedList(t)
		l.Add(100)
		l.Add(200)
		l.Add(300)

		l.Reverse()

		if l.head == nil || l.head.val != 300 {
			t.Fatalf("Novo head deveria ser 300, obtido %v", l.head)
		}
		if l.tail == nil || l.tail.val != 100 {
			t.Fatalf("Novo tail deveria ser 100, obtido %v", l.tail)
		}
		if l.head.prev != nil {
			t.Errorf("O prev do novo head deveria ser nil")
		}
		if l.tail.next != nil {
			t.Errorf("O next do novo tail deveria ser nil")
		}
	})

	t.Run("Reverse Bidirectional Traversal", func(t *testing.T) {
		l := newDoublyLinkedList(t)
		l.Add(10)
		l.Add(20)
		l.Add(30)
		l.Add(40)

		l.Reverse()

		// Travessia Forward (usando next a partir do head)
		curr := l.head
		expectedForward := []int{40, 30, 20, 10}
		for i, exp := range expectedForward {
			if curr == nil {
				t.Fatalf("Travessia forward parou prematuramente na iteração %d", i)
			}
			if curr.val != exp {
				t.Errorf("Forward: esperado %d, obtido %d", exp, curr.val)
			}
			curr = curr.next
		}

		// Travessia Backward (usando prev a partir do tail)
		curr = l.tail
		expectedBackward := []int{10, 20, 30, 40}
		for i, exp := range expectedBackward {
			if curr == nil {
				t.Fatalf("Travessia backward parou prematuramente na iteração %d", i)
			}
			if curr.val != exp {
				t.Errorf("Backward: esperado %d, obtido %d", exp, curr.val)
			}
			curr = curr.prev
		}
	})
}
