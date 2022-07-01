package matrix32

import "fmt"
import "bitmatrix"

type BitMatrix struct {
	Bits [][]uint32
	size int
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func NewBitMatrix(size int) Matrix {

	bitsize := (size / 32) + minInt((size%32), 1)
	matrix := &BitMatrix{Bits: make([][]uint32, size)}
	matrix.size = size

	for i := range matrix.Bits {
		matrix.Bits[i] = make([]uint32, bitsize)
	}

	return matrix
}

func (b *BitMatrix) GetMatrix(size int) Matrix {
	return NewBitMatrix(size)
}

func (b *BitMatrix) SetIndex(i, j int, val bool) {
	i = i - 1
	j = j - 1

	if i < 0 || j > b.size-1 {
		panic("index out of range ")
	}

	bitvector := b.Bits[i]
	index := j / 32
	bitnumber := uint32(j % 32)

	if val == true {
		bitvector[index] |= (1 << bitnumber)
	} else {
		mask := uint32(^(1 << bitnumber))
		bitvector[index] &= mask
	}

}

func (b *BitMatrix) GetIndex(i, j int) bool {
	i = i - 1
	j = j - 1

	if i < 0 || j > b.size-1 {
		panic("index out of range ")
	}

	bitvector := b.Bits[i]
	index := j / 32
	bitnumber := j % 32
	val := bitvector[index] & (1 << bitnumber)

	return (val > 0)
}

func (b *BitMatrix) PrintMatrix() {
	for i := 1; i <= b.size; i++ {
		for j := 1; j <= b.size; j++ {
			if b.GetIndex(i, j) == true {
				fmt.Print("1")
			} else {
				fmt.Print("0")
			}
		}
		fmt.Println("")
	}
}

// currently only support squared size
func (A *BitMatrix) Multiply(B Matrix) Matrix {
	if A.size != B.GetDimensions() {
		panic("A and B not the same size")
	}

	C := NewBitMatrix(A.size)

	for i := 1; i <= A.size; i++ {
		sum := false
		for j := 1; j <= A.size; j++ {
			for k := 1; k <= A.size; k++ {
				sum = (sum || (A.GetIndex(i, k) && B.GetIndex(k, j)))
			}
			C.SetIndex(i, j, sum)
		}
	}

	return C
}

func (A *BitMatrix) Add(B Matrix) Matrix {
	if A.size != B.GetDimensions() {
		panic("A and B not the same size")
	}
	C := NewBitMatrix(A.size)

	for i := 1; i < A.size; i++ {
		for j := 1; j < A.size; j++ {
			C.SetIndex(i, j, A.GetIndex(i, j) || B.GetIndex(i, j))
		}
	}
	return C
}

func (b *BitMatrix) GetDimensions() int {
	return b.size
}

func (b *BitMatrix) GetSizeInBytes() int {
	return 4 * len(b.Bits) * len(b.Bits[0])
}
