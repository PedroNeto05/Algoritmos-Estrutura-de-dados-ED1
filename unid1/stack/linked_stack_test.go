package stack

import (
	"testing"
)

// Factory para facilitar a criação nos testes
func newLinkedStack(t *testing.T) *LinkedStack {
	t.Helper()
	return &LinkedStack{}
}

func TestLinkedStackImplementation(t *testing.T) {
	t.Run("Push and Size", func(t *testing.T) {
		s := newLinkedStack(t)

		if !s.IsEmpty() {
			t.Error("Pilha nova deveria estar vazia")
		}

		s.Push(10)
		s.Push(20)

		if s.Size() != 2 {
			t.Errorf("Tamanho esperado 2, obtido %d", s.Size())
		}
		if s.IsEmpty() {
			t.Error("Pilha não deveria estar vazia após inserções")
		}
	})

	t.Run("LIFO Property (Last In, First Out)", func(t *testing.T) {
		s := newLinkedStack(t)

		s.Push(1) // Primeiro a entrar (ficará no fundo)
		s.Push(2)
		s.Push(3) // Último a entrar (ficará no topo)

		expectedValues := []int{3, 2, 1}

		for _, exp := range expectedValues {
			val, err := s.Pop()
			if err != nil {
				t.Fatalf("Pop retornou erro inesperado: %v", err)
			}
			if val != exp {
				t.Errorf("Erro no LIFO: esperado %d, obtido %d", exp, val)
			}
		}
	})

	t.Run("Peek Behavior", func(t *testing.T) {
		s := newLinkedStack(t)
		s.Push(42)

		// Peek não deve remover o elemento, apenas olhar
		val, err := s.Peek()
		if err != nil {
			t.Fatalf("Peek retornou erro inesperado: %v", err)
		}
		if val != 42 {
			t.Errorf("Peek falhou: esperado 42, obtido %d", val)
		}

		if s.Size() != 1 {
			t.Errorf("Tamanho mudou após Peek! Esperado 1, obtido %d", s.Size())
		}

		// O topo real (top node) deve continuar existindo
		if s.top == nil || s.top.val != 42 {
			t.Errorf("O nó do topo foi perdido ou alterado após o Peek")
		}

		// Pop deve retornar o exato mesmo elemento que o Peek viu
		popVal, err := s.Pop()
		if err != nil {
			t.Fatalf("Pop retornou erro inesperado: %v", err)
		}
		if popVal != 42 {
			t.Errorf("Pop diferente do Peek: esperado 42, obtido %d", popVal)
		}
	})

	t.Run("Pop Empty Stack", func(t *testing.T) {
		s := newLinkedStack(t)

		if _, err := s.Pop(); err == nil {
			t.Error("Deveria retornar erro ao dar Pop em pilha vazia")
		}
	})

	t.Run("Peek Empty Stack", func(t *testing.T) {
		s := newLinkedStack(t)

		if _, err := s.Peek(); err == nil {
			t.Error("Deveria retornar erro ao dar Peek em pilha vazia")
		}
	})

	t.Run("Pop Until Empty and Push Again", func(t *testing.T) {
		s := newLinkedStack(t)

		s.Push(100)
		_, err := s.Pop() // Esvaziou a pilha
		if err != nil {
			t.Fatalf("Pop retornou erro inesperado ao esvaziar: %v", err)
		}

		if !s.IsEmpty() {
			t.Error("Pilha deveria estar vazia")
		}

		if s.top != nil {
			t.Error("O ponteiro 'top' deveria ser nil após esvaziar a pilha")
		}

		// Deve aceitar novos elementos normalmente com os ponteiros renovados
		s.Push(200)
		val, err := s.Peek()
		if err != nil {
			t.Fatalf("Peek retornou erro inesperado após reinserção: %v", err)
		}
		if val != 200 {
			t.Errorf("Esperava 200 após reinserção, obtido %d", val)
		}
	})

	t.Run("Stress Test (Memory & Pointers)", func(t *testing.T) {
		s := newLinkedStack(t)

		// Push 10000 elementos
		for i := 0; i < 10000; i++ {
			s.Push(i)
		}

		if s.Size() != 10000 {
			t.Fatalf("Esperado 10000 elementos, obtido %d", s.Size())
		}

		// Pop 5000 elementos
		for i := 0; i < 5000; i++ {
			if _, err := s.Pop(); err != nil {
				t.Fatalf("Erro inesperado no Pop durante stress test: %v", err)
			}
		}

		if s.Size() != 5000 {
			t.Errorf("Esperado 5000 elementos, obtido %d", s.Size())
		}

		// O topo agora deve ser 4999 (pois o último inserido foi 9999 e tiramos 5000 de cima)
		val, err := s.Peek()
		if err != nil {
			t.Fatalf("Peek retornou erro inesperado após stress test: %v", err)
		}
		if val != 4999 {
			t.Errorf("Topo incorreto após stress test: esperado 4999, obtido %d", val)
		}
	})
}
