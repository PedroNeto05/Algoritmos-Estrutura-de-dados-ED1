// Programa de demonstração das estruturas implementadas na Unidade 3:
// árvore AVL, Tabela de Dispersão (Map) e Heap Binária (PriorityQueue).
//
// Execução: go run . (a partir da pasta unid3)
package main

import (
	"fmt"

	"unid3/avl"
	"unid3/hash"
	"unid3/heap"
)

func main() {
	demoAVL()
	demoHash()
	demoHeap()
}

func demoAVL() {
	fmt.Println("==================== AVL ====================")

	// Insere os caracteres de "Pedro Nascimento" (sem repeticao).
	letras := "PEDRONASCIMT"
	root := avl.NewNode(int(letras[0]))
	for i := 1; i < len(letras); i++ {
		root = root.Add(int(letras[i]))
	}

	fmt.Printf("Raiz apos inserir %q: %c (altura=%d)\n",
		letras, rune(root.Value()), root.RootHeight())
	fmt.Print("Em ordem (alfabetica): ")
	for _, v := range root.InOrder() {
		fmt.Printf("%c ", rune(v))
	}
	fmt.Printf("\nE uma AVL valida? %v\n", root.IsAVL())

	// Demonstra remocao com rebalanceamento.
	root = root.Remove(int('A'))
	root = root.Remove(int('T'))
	fmt.Print("Apos remover A e T:    ")
	for _, v := range root.InOrder() {
		fmt.Printf("%c ", rune(v))
	}
	fmt.Printf("\nContinua AVL valida?   %v\n\n", root.IsAVL())

	// Exemplo numerico (questao 22): insercao de 3 dispara rotacao em 10.
	nums := []int{13, 10, 15, 5, 11, 16, 4, 8}
	t := avl.NewNode(nums[0])
	for _, v := range nums[1:] {
		t = t.Add(v)
	}
	t = t.Add(3)
	fmt.Printf("Questao 22 -> nova raiz = %d, em ordem = %v (AVL=%v)\n\n",
		t.Value(), t.InOrder(), t.IsAVL())
}

func demoHash() {
	fmt.Println("=========== Tabela de Dispersao (Map) ===========")

	var m hash.Map = hash.NewHashTable(8)
	m.Put("um", 1)
	m.Put("dois", 2)
	m.Put("tres", 3)
	m.Put("dois", 22) // atualiza chave existente

	fmt.Printf("Tamanho: %d\n", m.Size())
	for _, k := range []string{"um", "dois", "tres", "quatro"} {
		if v, ok := m.Get(k); ok {
			fmt.Printf("  %-7s -> %d\n", k, v)
		} else {
			fmt.Printf("  %-7s -> (ausente)\n", k)
		}
	}
	fmt.Printf("Contem \"tres\"? %v\n", m.Contains("tres"))
	m.Remove("tres")
	fmt.Printf("Apos remover \"tres\": contem? %v, tamanho=%d\n\n",
		m.Contains("tres"), m.Size())
}

func demoHeap() {
	fmt.Println("=========== Heap Binaria (PriorityQueue) ===========")

	caracteres := "PEDRONASCI" // 10 primeiros de "Pedro Nascimento"

	// Min-heap: drena em ordem crescente.
	var minH heap.PQ = heap.NewMinHeap()
	for _, c := range caracteres {
		minH.Add(int(c))
	}
	fmt.Print("Min-heap (Poll ate esvaziar): ")
	drain(minH)

	// Max-heap: drena em ordem decrescente.
	var maxH heap.PQ = heap.NewMaxHeap()
	for _, c := range caracteres {
		maxH.Add(int(c))
	}
	fmt.Print("Max-heap (Poll ate esvaziar): ")
	drain(maxH)

	// Operacoes pedidas nas questoes 6 e 7 sobre a min-heap.
	q := heap.NewMinHeap()
	for _, c := range caracteres {
		q.Add(int(c))
	}
	if top, ok := q.Peek(); ok {
		fmt.Printf("Min-heap Peek (maior prioridade): %c\n", rune(top))
	}
	q.Add(int('A'))
	q.Add(int('Z'))
	removed, _ := q.Poll()
	q.Remove(int('P'))
	fmt.Printf("Apos Add A, Add Z, Poll(=%c) e Remove(P): ", rune(removed))
	drain(q)
}

// drain remove e imprime todos os elementos da fila em ordem de prioridade.
func drain(pq heap.PQ) {
	for !pq.IsEmpty() {
		v, _ := pq.Poll()
		fmt.Printf("%c ", rune(v))
	}
	fmt.Println()
}
