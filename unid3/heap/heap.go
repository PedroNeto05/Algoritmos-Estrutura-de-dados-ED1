// Package heap implementa uma Fila de Prioridade (PriorityQueue) usando uma
// Heap Binária armazenada em um array (slice).
//
// A mesma estrutura serve para Heap de Mínimo e Heap de Máximo, controlada
// pelo campo isMax. Para um nó no índice i (base 0):
//   - pai      = (i-1)/2
//   - filho esq = 2*i + 1
//   - filho dir = 2*i + 2
package heap

// Questão 3 — Funções fundamentais do tipo PriorityQueue.
type PQ interface {
	Add(value int)         // insere um elemento
	Peek() (int, bool)     // consulta o de maior prioridade sem remover
	Poll() (int, bool)     // remove e retorna o de maior prioridade
	Remove(value int) bool // remove uma ocorrência de um valor
	Size() int             // quantidade de elementos
	IsEmpty() bool
}

// Questão 4 — Estrutura de dado BinaryHeap.
type BinaryHeap struct {
	data  []int
	isMax bool
}

// NewMinHeap cria uma heap de mínimo (raiz = menor elemento).
func NewMinHeap() *BinaryHeap { return &BinaryHeap{} }

// NewMaxHeap cria uma heap de máximo (raiz = maior elemento).
func NewMaxHeap() *BinaryHeap { return &BinaryHeap{isMax: true} }

// higherPriority diz se a tem prioridade maior que b conforme o tipo de heap.
func (h *BinaryHeap) higherPriority(a, b int) bool {
	if h.isMax {
		return a > b
	}
	return a < b
}

func (h *BinaryHeap) swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

// siftUp sobe o elemento da posição i enquanto tiver prioridade maior que o pai.
func (h *BinaryHeap) siftUp(i int) {
	for i > 0 {
		parent := (i - 1) / 2
		if h.higherPriority(h.data[i], h.data[parent]) {
			h.swap(i, parent)
			i = parent
		} else {
			break
		}
	}
}

// siftDown desce o elemento da posição i para restaurar a propriedade de heap.
func (h *BinaryHeap) siftDown(i int) {
	n := len(h.data)
	for {
		left := 2*i + 1
		right := 2*i + 2
		best := i
		if left < n && h.higherPriority(h.data[left], h.data[best]) {
			best = left
		}
		if right < n && h.higherPriority(h.data[right], h.data[best]) {
			best = right
		}
		if best == i {
			break
		}
		h.swap(i, best)
		i = best
	}
}

// Questão 5 — Implementação das funções da PQ.

// Add insere um elemento na heap.
func (h *BinaryHeap) Add(value int) {
	h.data = append(h.data, value)
	h.siftUp(len(h.data) - 1)
}

// Peek retorna o elemento de maior prioridade sem removê-lo.
func (h *BinaryHeap) Peek() (int, bool) {
	if len(h.data) == 0 {
		return 0, false
	}
	return h.data[0], true
}

// Poll remove e retorna o elemento de maior prioridade (raiz).
func (h *BinaryHeap) Poll() (int, bool) {
	if len(h.data) == 0 {
		return 0, false
	}
	top := h.data[0]
	last := len(h.data) - 1
	h.data[0] = h.data[last]
	h.data = h.data[:last]
	if len(h.data) > 0 {
		h.siftDown(0)
	}
	return top, true
}

// Remove remove uma ocorrência do valor informado; true se encontrou.
func (h *BinaryHeap) Remove(value int) bool {
	idx := -1
	for i, v := range h.data {
		if v == value {
			idx = i
			break
		}
	}
	if idx == -1 {
		return false
	}
	last := len(h.data) - 1
	h.data[idx] = h.data[last]
	h.data = h.data[:last]
	if idx < len(h.data) {
		// pode precisar subir ou descer dependendo do elemento que veio do fim
		h.siftDown(idx)
		h.siftUp(idx)
	}
	return true
}

// Size retorna a quantidade de elementos.
func (h *BinaryHeap) Size() int { return len(h.data) }

// IsEmpty indica se a heap está vazia.
func (h *BinaryHeap) IsEmpty() bool { return len(h.data) == 0 }
