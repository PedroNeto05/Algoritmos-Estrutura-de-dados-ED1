package deque

import (
	"testing"
)

// Factory para facilitar a criação nos testes
func newArrayDeque(t *testing.T, initialSize int) *ArrayDeque {
	t.Helper()
	d := &ArrayDeque{}
	if err := d.Init(initialSize); err != nil {
		t.Fatalf("Falha ao inicializar Deque com tamanho %d: %v", initialSize, err)
	}
	return d
}

func TestArrayDeque(t *testing.T) {
	t.Run("Initialization Invalid Size", func(t *testing.T) {
		d := &ArrayDeque{}
		if err := d.Init(0); err == nil {
			t.Error("Deveria retornar erro para tamanho 0")
		}
		if err := d.Init(-5); err == nil {
			t.Error("Deveria retornar erro para tamanho negativo")
		}
	})

	t.Run("Empty Bounds Constraints", func(t *testing.T) {
		d := newArrayDeque(t, 2)

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

	t.Run("Queue Behavior (EnqueueRear, DequeueFront)", func(t *testing.T) {
		d := newArrayDeque(t, 5)

		d.EnqueueRear(10)
		d.EnqueueRear(20)
		d.EnqueueRear(30)

		if d.Size() != 3 {
			t.Errorf("Tamanho esperado 3, obtido %d", d.Size())
		}

		expected := []int{10, 20, 30}
		for _, exp := range expected {
			val, err := d.DequeueFront()
			if err != nil {
				t.Fatalf("Erro inesperado: %v", err)
			}
			if val != exp {
				t.Errorf("FIFO falhou: esperado %d, obtido %d", exp, val)
			}
		}
	})

	t.Run("Stack Behavior (EnqueueFront, DequeueFront)", func(t *testing.T) {
		d := newArrayDeque(t, 5)

		d.EnqueueFront(10)
		d.EnqueueFront(20)
		d.EnqueueFront(30)

		expected := []int{30, 20, 10} // LIFO
		for _, exp := range expected {
			val, err := d.DequeueFront()
			if err != nil {
				t.Fatalf("Erro inesperado: %v", err)
			}
			if val != exp {
				t.Errorf("LIFO falhou: esperado %d, obtido %d", exp, val)
			}
		}
	})

	t.Run("Mixed Ends (EnqueueFront, DequeueRear)", func(t *testing.T) {
		d := newArrayDeque(t, 5)

		d.EnqueueFront(1) // Vai pro fundo
		d.EnqueueFront(2) // Fica no meio
		d.EnqueueFront(3) // Fica na frente

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

	t.Run("Circular Wrap Around - Forwards (Rear)", func(t *testing.T) {
		d := newArrayDeque(t, 3)

		d.EnqueueRear(10)
		d.EnqueueRear(20)
		d.EnqueueRear(30)

		if _, err := d.DequeueFront(); err != nil {
			t.Fatal(err)
		} // tira 10
		if _, err := d.DequeueFront(); err != nil {
			t.Fatal(err)
		} // tira 20

		d.EnqueueRear(40) // Dá a volta pra frente
		d.EnqueueRear(50)

		expected := []int{30, 40, 50}
		for _, exp := range expected {
			val, err := d.DequeueFront()
			if err != nil {
				t.Fatal(err)
			}
			if val != exp {
				t.Errorf("Wrap around Forwards falhou: esperado %d, obtido %d", exp, val)
			}
		}
	})

	t.Run("Circular Wrap Around - Backwards (Front)", func(t *testing.T) {
		d := newArrayDeque(t, 3)

		// O array tem índices 0, 1, 2.
		d.EnqueueFront(10) // Cai no índice 2 (dá a volta pra trás)
		d.EnqueueFront(20) // Cai no índice 1
		d.EnqueueFront(30) // Cai no índice 0

		expected := []int{30, 20, 10}
		for _, exp := range expected {
			val, err := d.DequeueFront()
			if err != nil {
				t.Fatal(err)
			}
			if val != exp {
				t.Errorf("Wrap around Backwards falhou: esperado %d, obtido %d", exp, val)
			}
		}
	})

	t.Run("Dynamic Resize (Sequential)", func(t *testing.T) {
		d := newArrayDeque(t, 2)

		for i := 1; i <= 5; i++ {
			d.EnqueueRear(i) // Força redimensionamento 2 vezes
		}

		if d.Size() != 5 {
			t.Fatalf("Tamanho esperado 5, obtido %d", d.Size())
		}

		for i := 1; i <= 5; i++ {
			val, _ := d.DequeueFront()
			if val != i {
				t.Errorf("Erro no resize sequencial: esperado %d, obtido %d", i, val)
			}
		}
	})

	t.Run("Dynamic Resize (Wrapped Around)", func(t *testing.T) {
		// Teste crítico: Redimensionar quando o head está na frente do rear
		d := newArrayDeque(t, 3)

		d.EnqueueRear(10)
		d.EnqueueRear(20)
		d.EnqueueRear(30)

		if _, err := d.DequeueFront(); err != nil {
			t.Fatal(err)
		}
		if _, err := d.DequeueFront(); err != nil {
			t.Fatal(err)
		}

		d.EnqueueRear(40)
		d.EnqueueRear(50)
		// Fila: [30, 40, 50], head no meio, rear no começo.

		// Vai estourar a capacidade e forçar o unwrap
		d.EnqueueFront(60) // Deque agora é [60, 30, 40, 50]

		if d.Size() != 4 {
			t.Fatalf("Tamanho incorreto após resize: esperado 4, obtido %d", d.Size())
		}

		expected := []int{60, 30, 40, 50}
		for _, exp := range expected {
			val, err := d.DequeueFront()
			if err != nil {
				t.Fatal(err)
			}
			if val != exp {
				t.Errorf("Falha no resize em wrap-around: esperado %d, obtido %d", exp, val)
			}
		}
	})

	t.Run("Peek Methods (Front and Rear)", func(t *testing.T) {
		d := newArrayDeque(t, 5)

		d.EnqueueRear(100)
		d.EnqueueFront(50)
		d.EnqueueRear(200)
		// Estado atual: Frente [50, 100, 200] Fundo

		f, err := d.Front()
		if err != nil || f != 50 {
			t.Errorf("Front falhou: esperado 50, obtido %d", f)
		}

		r, err := d.Rear()
		if err != nil || r != 200 {
			t.Errorf("Rear falhou: esperado 200, obtido %d", r)
		}

		if d.Size() != 3 {
			t.Errorf("Peek alterou o tamanho: esperado 3, obtido %d", d.Size())
		}
	})
}
