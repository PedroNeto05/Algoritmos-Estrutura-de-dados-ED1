package bst

import "fmt"

type ITree interface {
	Add(value int)
	Search(value int) bool
	Min() int
	Max() int
	PrintPre()
	PrintIn()
	PrintPos()
	PrintLevels()
	NumPar() int
	Height() int
	Remove(value int) *BstNode
}
type BstNode struct {
	left  *BstNode
	value int
	right *BstNode
}

func (b *BstNode) Add(value int) {
	node := &BstNode{
		value: value,
	}

	if value <= b.value {
		if b.left != nil {
			b.left.Add(value)
		} else {
			b.left = node
		}
		return
	} else {
		if b.right != nil {
			b.right.Add(value)
		} else {
			b.right = node
		}
		return
	}
}

func (b *BstNode) Search(value int) bool {
	if b == nil {
		return false
	}
	if value < b.value {
		if b.left == nil {
			return false
		}
		return b.left.Search(value)
	} else if value > b.value {
		if b.right == nil {
			return false
		}
		return b.right.Search(value)
	}

	return true
}

func (b *BstNode) Min() int {
	if b == nil {
		return 0
	}
	if b.left != nil {
		return b.left.Min()
	}
	return b.value
}

func (b *BstNode) Max() int {
	if b == nil {
		return 0
	}
	if b.right != nil {
		return b.right.Max()
	}
	return b.value
}

func (b *BstNode) PrintPre() {
	if b == nil {
		return
	}

	fmt.Printf("%v ", b.value)
	if b.left != nil {
		b.left.PrintPre()
	}
	if b.right != nil {
		b.right.PrintPre()
	}
} // raiz-esquerda-direita
func (b *BstNode) PrintIn() {
	if b == nil {
		return
	}

	if b.left != nil {
		b.left.PrintIn()
	}

	fmt.Printf("%v ", b.value)

	if b.right != nil {
		b.right.PrintIn()
	}
} // esquerda-raiz-direita
func (b *BstNode) PrintPos() {
	if b == nil {
		return
	}

	if b.left != nil {
		b.left.PrintPos()
	}

	if b.right != nil {
		b.right.PrintPos()
	}

	fmt.Printf("%v ", b.value)
} // esquerda-direita-raiz
func (b *BstNode) PrintLevels() {
	if b == nil {
		return
	}

	fila := []*BstNode{b}

	for len(fila) > 0 {
		atual := fila[0]
		fila = fila[1:]

		fmt.Printf("%v ", atual)
	}
}

func (b *BstNode) NumPar() int {
	if b == nil {
		return 0
	}
	if b.value%2 == 0 {
		return 1 + b.left.NumPar() + b.right.NumPar()
	} else {
		return 0 + b.left.NumPar() + b.right.NumPar()
	}
}

func (b *BstNode) Height() int {
	if b == nil {
		return -1
	}

	alturaEsquerda := b.left.Height()
	alturaDireita := b.right.Height()

	if alturaDireita > alturaEsquerda {
		return alturaDireita + 1
	}

	return alturaEsquerda + 1
}

func (b *BstNode) Remove(value int) *BstNode {
	if b == nil {
		return nil
	}

	if value < b.value {
		b.left = b.left.Remove(value)
		return b
	}

	if value > b.value {
		b.right = b.right.Remove(value)
		return b
	}

	if b.left == nil && b.right == nil {
		return nil
	}

	if b.left == nil {
		return b.right
	}

	if b.right == nil {
		return b.left
	}

	sucessor := b.right
	for sucessor.left != nil {
		sucessor = sucessor.left
	}

	b.value = sucessor.value

	b.right = b.right.Remove(sucessor.value)

	return b
}
