// Package avl implementa uma árvore AVL (BST auto-balanceada).
//
// Convenção de fator de balanceamento (bf) usada em toda a unidade,
// igual à do material do professor (questão 12):
//
//	bf = altura(subárvore direita) − altura(subárvore esquerda)
//
// Logo: bf < 0  => árvore pesada à ESQUERDA
//
//	bf > 0  => árvore pesada à DIREITA
//
// Altura: folha tem altura 0; uma subárvore vazia (nil) tem altura -1.
package avl

// BstNode é o nó da árvore (estrutura dada no enunciado).
type BstNode struct {
	left   *BstNode
	value  int
	height int
	bf     int
	right  *BstNode
}

// NewNode cria um novo nó folha já com as propriedades inicializadas.
func NewNode(value int) *BstNode {
	return &BstNode{value: value, height: 0, bf: 0}
}

// height retorna a altura de um nó tratando o caso nil (-1).
func height(n *BstNode) int {
	if n == nil {
		return -1
	}
	return n.height
}

// ---------------------------------------------------------------------------
// Questão 5 / 8 / 11 — Rotações
//
// As rotações já incluem as chamadas a UpdateProperties() (questão 11).
// Atualiza-se primeiro o nó que DESCEU (antiga raiz) e depois o que SUBIU
// (nova raiz), pois a altura/bf de quem sobe depende dos filhos já corrigidos.
// ---------------------------------------------------------------------------

// RotRight rotaciona a subárvore à direita.
func (root *BstNode) RotRight() *BstNode {
	newRoot := root.left
	root.left = newRoot.right
	newRoot.right = root

	root.UpdateProperties()    // antiga raiz (desceu) primeiro
	newRoot.UpdateProperties() // nova raiz depois
	return newRoot
}

// RotLeft rotaciona a subárvore à esquerda.
func (root *BstNode) RotLeft() *BstNode {
	newRoot := root.right
	root.right = newRoot.left
	newRoot.left = root

	root.UpdateProperties()
	newRoot.UpdateProperties()
	return newRoot
}

// ---------------------------------------------------------------------------
// Questão 10 — Atualização de propriedades (altura e fator de balanço)
// ---------------------------------------------------------------------------

// UpdateProperties recalcula altura e bf do nó a partir de seus filhos.
func (root *BstNode) UpdateProperties() *BstNode {
	lh := height(root.left)
	rh := height(root.right)

	if lh > rh {
		root.height = lh + 1
	} else {
		root.height = rh + 1
	}
	root.bf = rh - lh
	return root
}

// ---------------------------------------------------------------------------
// Questão 13 — Funções de rebalanceamento (6 casos)
// ---------------------------------------------------------------------------

// esquerda-esquerda: rotação simples à direita.
func (root *BstNode) RebalanceLeftLeft() *BstNode {
	return root.RotRight()
}

// esquerda-neutro: rotação simples à direita (ocorre em remoções).
func (root *BstNode) RebalanceLeftNeutral() *BstNode {
	return root.RotRight()
}

// esquerda-direita: rotação dupla -> esquerda no filho esquerdo, direita na raiz.
func (root *BstNode) RebalanceLeftRight() *BstNode {
	root.left = root.left.RotLeft()
	return root.RotRight()
}

// direita-direita: rotação simples à esquerda.
func (root *BstNode) RebalanceRightRight() *BstNode {
	return root.RotLeft()
}

// direita-neutro: rotação simples à esquerda (ocorre em remoções).
func (root *BstNode) RebalanceRightNeutral() *BstNode {
	return root.RotLeft()
}

// direita-esquerda: rotação dupla -> direita no filho direito, esquerda na raiz.
func (root *BstNode) RebalanceRightLeft() *BstNode {
	root.right = root.right.RotRight()
	return root.RotLeft()
}

// ---------------------------------------------------------------------------
// Questão 15 — Rebalanceamento genérico
// ---------------------------------------------------------------------------

// Rebalance atualiza as propriedades do nó e, se ele estiver desbalanceado
// (|bf| > 1), aplica o caso de rebalanceamento adequado.
func (root *BstNode) Rebalance() *BstNode {
	root.UpdateProperties()

	if root.bf < -1 { // pesada à esquerda
		if root.left.bf < 0 {
			return root.RebalanceLeftLeft()
		} else if root.left.bf > 0 {
			return root.RebalanceLeftRight()
		}
		return root.RebalanceLeftNeutral()
	}

	if root.bf > 1 { // pesada à direita
		if root.right.bf > 0 {
			return root.RebalanceRightRight()
		} else if root.right.bf < 0 {
			return root.RebalanceRightLeft()
		}
		return root.RebalanceRightNeutral()
	}

	return root
}

// ---------------------------------------------------------------------------
// Questão 16 — Inserção em AVL
// ---------------------------------------------------------------------------

// Add insere um valor e devolve a (possivelmente nova) raiz da subárvore já
// rebalanceada.
func (root *BstNode) Add(value int) *BstNode {
	if value <= root.value { // inserção à esquerda
		if root.left == nil {
			root.left = NewNode(value)
		} else {
			root.left = root.left.Add(value)
		}
	} else { // inserção à direita
		if root.right == nil {
			root.right = NewNode(value)
		} else {
			root.right = root.right.Add(value)
		}
	}
	return root.Rebalance()
}

// ---------------------------------------------------------------------------
// Questão 18 — Remoção em AVL
// ---------------------------------------------------------------------------

// Max retorna o maior valor da subárvore (mais à direita).
func (root *BstNode) Max() int {
	if root.right == nil {
		return root.value
	}
	return root.right.Max()
}

// Remove remove um valor e devolve a (possivelmente nova) raiz da subárvore,
// rebalanceando no caminho de volta.
func (root *BstNode) Remove(value int) *BstNode {
	if root == nil {
		return nil
	}

	if value < root.value {
		root.left = root.left.Remove(value)
	} else if value > root.value {
		root.right = root.right.Remove(value)
	} else { // encontramos o nó a ser removido
		if root.left == nil && root.right == nil { // folha
			return nil
		} else if root.left != nil && root.right == nil { // 1 filho (esq)
			return root.left
		} else if root.left == nil && root.right != nil { // 1 filho (dir)
			return root.right
		} else { // 2 filhos
			maxEsq := root.left.Max()
			root.value = maxEsq
			root.left = root.left.Remove(maxEsq)
		}
	}
	return root.Rebalance()
}

// ---------------------------------------------------------------------------
// Questão 20 — Verifica se a árvore é AVL
// ---------------------------------------------------------------------------

// realHeight recalcula a altura real (sem confiar no campo height).
func realHeight(n *BstNode) int {
	if n == nil {
		return -1
	}
	l := realHeight(n.left)
	r := realHeight(n.right)
	if l > r {
		return l + 1
	}
	return r + 1
}

// InOrder retorna os valores da árvore em ordem (crescente).
func (root *BstNode) InOrder() []int {
	if root == nil {
		return nil
	}
	out := root.left.InOrder()
	out = append(out, root.value)
	out = append(out, root.right.InOrder()...)
	return out
}

// Value retorna o valor armazenado no nó (útil para inspecionar a raiz).
func (root *BstNode) Value() int { return root.value }

// Height retorna a altura armazenada no nó.
func (root *BstNode) RootHeight() int { return root.height }

// IsAVL retorna true se, para todo nó, |bf| <= 1.
func (root *BstNode) IsAVL() bool {
	if root == nil {
		return true
	}
	bf := realHeight(root.right) - realHeight(root.left)
	if bf < -1 || bf > 1 {
		return false
	}
	return root.left.IsAVL() && root.right.IsAVL()
}
