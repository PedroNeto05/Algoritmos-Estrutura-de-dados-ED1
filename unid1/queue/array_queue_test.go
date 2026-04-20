package queue

import (
	"testing"
)

// Factory para facilitar a criação nos testes
func newArrayQueue(t *testing.T, initialSize int) *ArrayQueue {
	t.Helper()
	q := &ArrayQueue{}
	if err := q.Init(initialSize); err != nil {
		t.Fatalf("Falha ao inicializar fila com tamanho %d: %v", initialSize, err)
	}
	return q
}

func TestArrayQueueImplementation(t *testing.T) {
	t.Run("Initialization Invalid Size", func(t *testing.T) {
		q := &ArrayQueue{}
		if err := q.Init(0); err == nil {
			t.Error("Deveria retornar erro para tamanho 0")
		}
		if err := q.Init(-5); err == nil {
			t.Error("Deveria retornar erro para tamanho negativo")
		}
	})

	t.Run("Basic Enqueue and Size", func(t *testing.T) {
		q := newArrayQueue(t, 5)

		if !q.IsEmpty() {
			t.Error("Fila nova deveria estar vazia")
		}

		q.Enqueue(10)
		q.Enqueue(20)

		if q.Size() != 2 {
			t.Errorf("Tamanho esperado 2, obtido %d", q.Size())
		}
		if q.IsEmpty() {
			t.Error("Fila não deveria estar vazia após inserções")
		}
	})

	t.Run("FIFO Property (First In, First Out)", func(t *testing.T) {
		q := newArrayQueue(t, 3)

		q.Enqueue(1) // Primeiro a entrar
		q.Enqueue(2)
		q.Enqueue(3) // Último a entrar

		expectedValues := []int{1, 2, 3}

		for _, exp := range expectedValues {
			val, err := q.Dequeue()
			if err != nil {
				t.Fatalf("Dequeue retornou erro inesperado: %v", err)
			}
			if val != exp {
				t.Errorf("Erro no FIFO: esperado %d, obtido %d", exp, val)
			}
		}
	})

	t.Run("Front Behavior", func(t *testing.T) {
		q := newArrayQueue(t, 2)
		q.Enqueue(42)

		// Front não deve remover o elemento
		val, err := q.Front()
		if err != nil {
			t.Fatalf("Front retornou erro inesperado: %v", err)
		}
		if val != 42 {
			t.Errorf("Front falhou: esperado 42, obtido %d", val)
		}

		if q.Size() != 1 {
			t.Errorf("Tamanho mudou após Front! Esperado 1, obtido %d", q.Size())
		}

		// Dequeue deve retornar o mesmo elemento que o Front viu
		deqVal, err := q.Dequeue()
		if err != nil {
			t.Fatalf("Dequeue retornou erro inesperado: %v", err)
		}
		if deqVal != 42 {
			t.Errorf("Dequeue diferente do Front: esperado 42, obtido %d", deqVal)
		}
	})

	t.Run("Empty Queue Constraints", func(t *testing.T) {
		q := newArrayQueue(t, 2)

		if _, err := q.Dequeue(); err == nil {
			t.Error("Deveria retornar erro ao dar Dequeue em fila vazia")
		}
		if _, err := q.Front(); err == nil {
			t.Error("Deveria retornar erro ao dar Front em fila vazia")
		}
	})

	t.Run("Size Updates Correctly After Dequeue", func(t *testing.T) {
		q := newArrayQueue(t, 5)

		q.Enqueue(10)
		q.Enqueue(20)

		if q.Size() != 2 {
			t.Fatalf("Tamanho inicial incorreto: esperado 2, obtido %d", q.Size())
		}

		_, err := q.Dequeue()
		if err != nil {
			t.Fatalf("Erro inesperado no Dequeue: %v", err)
		}

		if q.Size() != 1 {
			t.Errorf("Tamanho não foi atualizado após Dequeue: esperado 1, obtido %d", q.Size())
		}
	})

	t.Run("Circular Behavior (No Resize)", func(t *testing.T) {
		// Fila com capacidade 3
		q := newArrayQueue(t, 3)

		// Enche a fila
		q.Enqueue(1)
		q.Enqueue(2)
		q.Enqueue(3)

		// Remove 2 (head anda pra frente, libera espaço no começo do array)
		if _, err := q.Dequeue(); err != nil {
			t.Fatalf("Erro inesperado no Dequeue 1: %v", err)
		}
		if _, err := q.Dequeue(); err != nil {
			t.Fatalf("Erro inesperado no Dequeue 2: %v", err)
		}

		// Adiciona mais 2 (tail deve dar a volta no array)
		q.Enqueue(4)
		q.Enqueue(5)

		// Verifica se a ordem FIFO se mantém perfeita
		expectedValues := []int{3, 4, 5}
		for _, exp := range expectedValues {
			val, err := q.Dequeue()
			if err != nil {
				t.Fatalf("Dequeue falhou: %v", err)
			}
			if val != exp {
				t.Errorf("Erro na circularidade: esperado %d, obtido %d", exp, val)
			}
		}
	})

	t.Run("Dynamic Resize (Sequential)", func(t *testing.T) {
		// Inicializa com tamanho pequeno
		q := newArrayQueue(t, 2)

		// Estoura a capacidade (força o duplicateQueue)
		for i := 1; i <= 5; i++ {
			q.Enqueue(i)
		}

		if q.Size() != 5 {
			t.Errorf("Tamanho esperado 5, obtido %d", q.Size())
		}

		expected := []int{1, 2, 3, 4, 5}
		for _, exp := range expected {
			val, err := q.Dequeue()
			if err != nil {
				t.Fatalf("Dequeue falhou após resize: %v", err)
			}
			if val != exp {
				t.Errorf("Após resize, valor esperado %d, obtido %d", exp, val)
			}
		}
	})

	t.Run("Dynamic Resize (While Wrapped Around)", func(t *testing.T) {
		// Este é o teste de fogo para arrays circulares dinâmicos
		q := newArrayQueue(t, 3)

		// Adiciona 3 e remove 2. (Array interno tem espaço vazio no começo)
		q.Enqueue(10)
		q.Enqueue(20)
		q.Enqueue(30)

		if _, err := q.Dequeue(); err != nil {
			t.Fatalf("Erro inesperado: %v", err)
		} // tira o 10
		if _, err := q.Dequeue(); err != nil {
			t.Fatalf("Erro inesperado: %v", err)
		} // tira o 20
		// A fila atual tem o [30]

		// Adiciona mais itens para fazer o tail dar a volta (wrap around)
		q.Enqueue(40)
		q.Enqueue(50)
		// A fila tem [30, 40, 50]. Head aponta pro meio, tail aponta pro começo.

		// O próximo Enqueue VAI forçar o resize enquanto os ponteiros estão invertidos
		q.Enqueue(60)

		// Valida se a ordem se manteve correta após esticar o array que estava "dobrado"
		expected := []int{30, 40, 50, 60}
		for _, exp := range expected {
			val, err := q.Dequeue()
			if err != nil {
				t.Fatalf("Dequeue falhou após resize complexo: %v", err)
			}
			if val != exp {
				t.Errorf("Após resize complexo, valor esperado %d, obtido %d", exp, val)
			}
		}
	})
	t.Run("SizeAlt Calculation", func(t *testing.T) {
		// Inicializamos com 5
		q := newArrayQueue(t, 5)

		// 1. Fila Vazia
		if got := q.SizeAlt(); got != 0 {
			t.Errorf("Fila vazia: SizeAlt esperado 0, obtido %d", got)
		}

		// 2. Fila com 1 elemento
		q.Enqueue(10)
		if got := q.SizeAlt(); got != 1 {
			t.Errorf("1 elemento: SizeAlt esperado 1, obtido %d", got)
		}

		// 3. Fila com múltiplos elementos (linear, sem dar a volta)
		q.Enqueue(20)
		q.Enqueue(30)
		if got := q.SizeAlt(); got != 3 {
			t.Errorf("3 elementos: SizeAlt esperado 3, obtido %d", got)
		}

		// 4. Forçar o Wrap-Around (Head e Tail dão a volta)
		if _, err := q.Dequeue(); err != nil { // remove 10 (head vai para 1)
			t.Fatalf("Erro inesperado no Dequeue 1: %v", err)
		}
		if _, err := q.Dequeue(); err != nil { // remove 20 (head vai para 2)
			t.Fatalf("Erro inesperado no Dequeue 2: %v", err)
		}

		// Neste momento tem 1 elemento [30]. Vamos adicionar mais 3.
		q.Enqueue(40)
		q.Enqueue(50)
		q.Enqueue(60)
		// A fila tem 4 elementos [30, 40, 50, 60].
		// O tail deu a volta e deve estar num índice menor que o head!

		if got := q.SizeAlt(); got != 4 {
			t.Errorf("Wrap around: SizeAlt esperado 4, obtido %d", got)
		}
	})
}
