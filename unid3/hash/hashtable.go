// Package hash implementa uma Tabela de Dispersão (Hash Table) que realiza o
// TAD Map, usando encadeamento separado (separate chaining) para tratar colisões.
//
// Estrutura de dados básica (questão 1): um array (slice) de buckets, indexado
// pelo valor de hash da chave.
package hash

// Questão 3 — Funções fundamentais do tipo Map.
type Map interface {
	Put(key string, value int)  // insere/atualiza um par chave-valor
	Get(key string) (int, bool) // retorna o valor e se a chave existe
	Remove(key string) bool     // remove uma chave; true se existia
	Contains(key string) bool   // true se a chave está presente
	Size() int                  // quantidade de pares armazenados
	IsEmpty() bool
}

// entry é um nó da lista ligada de um bucket (resolução de colisão).
type entry struct {
	key   string
	value int
	next  *entry
}

// Questão 4 — Estrutura de dado HashTable.
type HashTable struct {
	buckets []*entry
	size    int
}

// NewHashTable cria uma tabela com a capacidade (número de buckets) informada.
func NewHashTable(capacity int) *HashTable {
	if capacity <= 0 {
		capacity = 16
	}
	return &HashTable{buckets: make([]*entry, capacity)}
}

// hash mapeia uma chave (string) para um índice do array (questão 2).
// Usa a clássica função polinomial base 31 e depois aplica módulo.
func (h *HashTable) hash(key string) int {
	sum := 0
	for _, c := range key {
		sum = sum*31 + int(c)
	}
	if sum < 0 {
		sum = -sum
	}
	return sum % len(h.buckets)
}

// Questão 6 — Implementação das funções do Map.

// Put insere um novo par ou atualiza o valor de uma chave existente.
func (h *HashTable) Put(key string, value int) {
	idx := h.hash(key)
	for e := h.buckets[idx]; e != nil; e = e.next {
		if e.key == key {
			e.value = value // chave já existe: atualiza
			return
		}
	}
	// insere no início da lista do bucket
	h.buckets[idx] = &entry{key: key, value: value, next: h.buckets[idx]}
	h.size++
}

// Get retorna o valor associado à chave e um bool indicando se foi encontrada.
func (h *HashTable) Get(key string) (int, bool) {
	idx := h.hash(key)
	for e := h.buckets[idx]; e != nil; e = e.next {
		if e.key == key {
			return e.value, true
		}
	}
	return 0, false
}

// Remove remove a chave; retorna true se ela existia.
func (h *HashTable) Remove(key string) bool {
	idx := h.hash(key)
	var prev *entry
	for e := h.buckets[idx]; e != nil; e = e.next {
		if e.key == key {
			if prev == nil {
				h.buckets[idx] = e.next
			} else {
				prev.next = e.next
			}
			h.size--
			return true
		}
		prev = e
	}
	return false
}

// Contains indica se a chave está presente.
func (h *HashTable) Contains(key string) bool {
	_, ok := h.Get(key)
	return ok
}

// Size retorna a quantidade de pares armazenados.
func (h *HashTable) Size() int {
	return h.size
}

// IsEmpty indica se a tabela está vazia.
func (h *HashTable) IsEmpty() bool {
	return h.size == 0
}
