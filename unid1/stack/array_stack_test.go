package stack

import (
	"testing"
)

// Factory para facilitar a criação nos testes
func newArrayStack(t *testing.T, initialSize int) *ArrayStack {
	t.Helper()
	s := &ArrayStack{}
	if err := s.Init(initialSize); err != nil {
		t.Fatalf("Falha ao inicializar pilha com tamanho %d: %v", initialSize, err)
	}
	return s
}

func TestArrayStackImplementation(t *testing.T) {
	t.Run("Initialization Invalid Size", func(t *testing.T) {
		s := &ArrayStack{}
		if err := s.Init(0); err == nil {
			t.Error("Deveria retornar erro para tamanho 0")
		}
		if err := s.Init(-5); err == nil {
			t.Error("Deveria retornar erro para tamanho negativo")
		}
	})

	t.Run("Push and Size", func(t *testing.T) {
		s := newArrayStack(t, 5)

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
		s := newArrayStack(t, 3)

		s.Push(1) // Primeiro a entrar
		s.Push(2)
		s.Push(3) // Último a entrar

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

	t.Run("Dynamic Resize (duplicateStack)", func(t *testing.T) {
		// Inicializa com tamanho 2
		s := newArrayStack(t, 2)

		// Insere 5 elementos (vai forçar o duplicateStack a rodar)
		for i := 1; i <= 5; i++ {
			s.Push(i)
		}

		if s.Size() != 5 {
			t.Errorf("Tamanho esperado 5, obtido %d", s.Size())
		}

		// Valida se os elementos continuam na ordem correta após o resize
		expected := []int{5, 4, 3, 2, 1}
		for _, exp := range expected {
			val, err := s.Pop()
			if err != nil {
				t.Fatalf("Pop retornou erro inesperado após resize: %v", err)
			}
			if val != exp {
				t.Errorf("Após resize, valor esperado %d, obtido %d", exp, val)
			}
		}
	})

	t.Run("Peek Behavior", func(t *testing.T) {
		s := newArrayStack(t, 2)
		s.Push(42)

		// Peek não deve remover o elemento
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

		// Pop deve retornar o mesmo elemento que o Peek viu
		popVal, err := s.Pop()
		if err != nil {
			t.Fatalf("Pop retornou erro inesperado: %v", err)
		}
		if popVal != 42 {
			t.Errorf("Pop diferente do Peek: esperado 42, obtido %d", popVal)
		}
	})

	t.Run("Pop Empty Stack", func(t *testing.T) {
		s := newArrayStack(t, 2)

		if _, err := s.Pop(); err == nil {
			t.Error("Deveria retornar erro ao dar Pop em pilha vazia")
		}
	})

	t.Run("Peek Empty Stack", func(t *testing.T) {
		s := newArrayStack(t, 2)

		if _, err := s.Peek(); err == nil {
			t.Error("Deveria retornar erro ao dar Peek em pilha vazia")
		}
	})

	t.Run("Pop Until Empty and Push Again", func(t *testing.T) {
		s := newArrayStack(t, 2)

		s.Push(100)
		_, err := s.Pop() // Esvaziou
		if err != nil {
			t.Fatalf("Pop retornou erro inesperado ao esvaziar: %v", err)
		}

		if !s.IsEmpty() {
			t.Error("Pilha deveria estar vazia")
		}

		// Deve aceitar novos elementos normalmente
		s.Push(200)
		val, err := s.Peek()
		if err != nil {
			t.Fatalf("Peek retornou erro inesperado após reinserção: %v", err)
		}
		if val != 200 {
			t.Errorf("Esperava 200 após reinserção, obtido %d", val)
		}
	})

	t.Run("Stress Test", func(t *testing.T) {
		s := newArrayStack(t, 10)

		// Push 1000 elementos
		for i := 0; i < 1000; i++ {
			s.Push(i)
		}

		if s.Size() != 1000 {
			t.Fatalf("Esperado 1000 elementos, obtido %d", s.Size())
		}

		// Pop 500 elementos com tratamento de erro em massa
		for i := 0; i < 500; i++ {
			if _, err := s.Pop(); err != nil {
				t.Fatalf("Erro inesperado no Pop durante stress test: %v", err)
			}
		}

		if s.Size() != 500 {
			t.Errorf("Esperado 500 elementos, obtido %d", s.Size())
		}

		// O topo agora deve ser 499 (pois o último inserido foi 999 e tiramos os 500 de cima)
		val, err := s.Peek()
		if err != nil {
			t.Fatalf("Peek retornou erro inesperado após stress test: %v", err)
		}
		if val != 499 {
			t.Errorf("Topo incorreto após stress test: esperado 499, obtido %d", val)
		}
	})
}
