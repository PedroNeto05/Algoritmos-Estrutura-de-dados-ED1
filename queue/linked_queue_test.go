package queue

import (
	"testing"
)

// factory para a LinkedQueue
func newLinkedQueue(t *testing.T) *LinkedQueue {
	t.Helper()
	return &LinkedQueue{}
}

func TestLinkedQueueImplementation(t *testing.T) {
	t.Run("Initial State", func(t *testing.T) {
		q := newLinkedQueue(t)

		if !q.IsEmpty() {
			t.Error("Esperado que a fila recém-criada estivesse vazia")
		}

		if q.Size() != 0 {
			t.Errorf("Esperado tamanho 0, obtido %d", q.Size())
		}
	})

	t.Run("Enqueue and Size", func(t *testing.T) {
		q := newLinkedQueue(t)

		q.Enqueue(10)
		q.Enqueue(20)

		if q.Size() != 2 {
			t.Errorf("Esperado tamanho 2, obtido %d", q.Size())
		}

		if q.IsEmpty() {
			t.Error("A fila não deveria estar vazia após inserções")
		}
	})

	t.Run("Enqueue Multiple Elements", func(t *testing.T) {
		q := newLinkedQueue(t)

		for i := range 50 {
			q.Enqueue(i)
		}

		if q.Size() != 50 {
			t.Errorf("Esperado tamanho 50, obtido %d", q.Size())
		}
	})

	t.Run("Front Valid", func(t *testing.T) {
		q := newLinkedQueue(t)
		q.Enqueue(100)
		q.Enqueue(200)

		val, err := q.Front()
		if err != nil || val != 100 {
			t.Errorf("Front falhou: esperado 100, obtido val=%d err=%v", val, err)
		}

		// Garantir que Front não remove o elemento
		if q.Size() != 2 {
			t.Errorf("Esperado tamanho 2 após Front, obtido %d", q.Size())
		}
	})

	t.Run("Front Empty Queue", func(t *testing.T) {
		q := newLinkedQueue(t)

		if _, err := q.Front(); err == nil {
			t.Error("Esperava erro ao chamar Front em fila vazia")
		}
	})

	t.Run("Dequeue Valid", func(t *testing.T) {
		q := newLinkedQueue(t)
		q.Enqueue(10)
		q.Enqueue(20)

		val, err := q.Dequeue()
		if err != nil || val != 10 {
			t.Errorf("Dequeue falhou: esperado 10, obtido val=%d err=%v", val, err)
		}

		if q.Size() != 1 {
			t.Errorf("Esperado tamanho 1 após Dequeue, obtido %d", q.Size())
		}

		val2, err2 := q.Front()
		if err2 != nil || val2 != 20 {
			t.Errorf("Front falhou após Dequeue: esperado 20, obtido val=%d err=%v", val2, err2)
		}
	})

	t.Run("Dequeue Empty Queue", func(t *testing.T) {
		q := newLinkedQueue(t)

		if _, err := q.Dequeue(); err == nil {
			t.Error("Esperava erro ao chamar Dequeue em fila vazia")
		}
	})

	t.Run("Dequeue Until Empty", func(t *testing.T) {
		q := newLinkedQueue(t)

		q.Enqueue(1)
		q.Enqueue(2)

		if _, err := q.Dequeue(); err != nil {
			t.Fatalf("Erro no Dequeue: %v", err)
		}

		if _, err := q.Dequeue(); err != nil {
			t.Fatalf("Erro no Dequeue: %v", err)
		}

		if !q.IsEmpty() {
			t.Error("A fila deveria estar vazia após remover todos os elementos")
		}

		if q.Size() != 0 {
			t.Errorf("Esperado tamanho 0, obtido %d", q.Size())
		}

		// Tentar remover mais um deve dar erro
		if _, err := q.Dequeue(); err == nil {
			t.Error("Esperava erro ao remover de fila que acabou de ser esvaziada")
		}
	})

	t.Run("Stress Mixed Operations", func(t *testing.T) {
		q := newLinkedQueue(t)

		// Adiciona 100 elementos
		for i := range 100 {
			q.Enqueue(i)
		}

		// Remove os 50 primeiros e verifica se a ordem FIFO foi mantida
		for i := range 50 {
			val, err := q.Dequeue()
			if err != nil || val != i {
				t.Fatalf("Erro no Dequeue esperado %d, obtido %d: %v", i, val, err)
			}
		}

		if q.Size() != 50 {
			t.Errorf("Esperado tamanho 50, obtido %d", q.Size())
		}

		// Adiciona mais 50 elementos (100 a 149)
		for i := 100; i < 150; i++ {
			q.Enqueue(i)
		}

		if q.Size() != 100 {
			t.Errorf("Esperado tamanho 100, obtido %d", q.Size())
		}

		// Verifica se o Front atual é 50 (o próximo da fila original que não foi removido)
		frontVal, _ := q.Front()
		if frontVal != 50 {
			t.Errorf("Esperado Front 50, obtido %d", frontVal)
		}
	})
}

