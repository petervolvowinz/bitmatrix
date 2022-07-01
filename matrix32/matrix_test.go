/******** Peter Winzell (c), 7/1/22 *********************************************/

package matrix32

import (
	"github.com/petervolvowinz/bitmatrix"
	"testing"
)

func TestPrintMatrix(T *testing.T) {
	A := NewBitMatrix(57)
	// A.PrintMatrix()

	for i := 1; i <= 57; i++ {
		for j := 1; j <= 57; j++ {
			A.SetIndex(i, j, true)
		}
	}

	val := A.GetIndex(1, 1)
	if val == false {
		T.Error("EXPECTED TRUE AT M(1,1,)")
	}

	val = A.GetIndex(57, 57)
	if val == false {
		T.Error("EXPECTED TRUE AT M(57,57)")
	}

	val = A.GetIndex(56, 56)
	if val == false {
		T.Error("EXPECTED TRUE AT M(56,56)")
	}

	A.PrintMatrix()
}

func testMatrices(C, RES bitmatrix.Matrix) bool {
	for i := 1; i <= 2; i++ {
		for j := 1; j <= 2; j++ {
			if C.GetIndex(i, j) != RES.GetIndex(i, j) {
				C.PrintMatrix()
				RES.PrintMatrix()
				return false
			}
		}
	}
	return true
}

func TestMultiplication(T *testing.T) {

	FacitMatrix := NewBitMatrix(2)
	A := NewBitMatrix(2)
	B := NewBitMatrix(2)

	FacitMatrix.SetIndex(1, 1, false)
	FacitMatrix.SetIndex(1, 2, true)
	FacitMatrix.SetIndex(2, 1, false)
	FacitMatrix.SetIndex(2, 2, true)

	A.SetIndex(1, 1, true)
	A.SetIndex(1, 2, false)
	A.SetIndex(2, 1, false)
	A.SetIndex(2, 2, true)

	B.SetIndex(1, 1, false)
	B.SetIndex(1, 2, true)
	B.SetIndex(2, 1, false)
	B.SetIndex(2, 2, true)

	C := A.Multiply(B)

	if testMatrices(C, FacitMatrix) == false {
		T.Error("error in matrix multiplication")
	}

	FacitMatrix.SetIndex(1, 1, false)
	FacitMatrix.SetIndex(1, 2, false)
	FacitMatrix.SetIndex(2, 1, false)
	FacitMatrix.SetIndex(2, 2, false)

	A.SetIndex(1, 1, false)
	A.SetIndex(1, 2, false)
	A.SetIndex(2, 1, false)
	A.SetIndex(2, 2, false)

	B.SetIndex(1, 1, true)
	B.SetIndex(1, 2, true)
	B.SetIndex(2, 1, true)
	B.SetIndex(2, 2, true)

	C = A.Multiply(B)

	if testMatrices(C, FacitMatrix) == false {
		T.Error("error in matrix multiplication")
	}

	FacitMatrix.SetIndex(1, 1, true)
	FacitMatrix.SetIndex(1, 2, true)
	FacitMatrix.SetIndex(2, 1, true)
	FacitMatrix.SetIndex(2, 2, true)

	A.SetIndex(1, 1, true)
	A.SetIndex(1, 2, false)
	A.SetIndex(2, 1, false)
	A.SetIndex(2, 2, true)

	B.SetIndex(1, 1, true)
	B.SetIndex(1, 2, true)
	B.SetIndex(2, 1, true)
	B.SetIndex(2, 2, true)

	C = A.Multiply(B)

	if testMatrices(C, FacitMatrix) == false {
		T.Error("error in matrix multiplication")
	}

}

func TestGetSize(T *testing.T) {
	A := NewBitMatrix(64)
	bytes := A.GetSizeInBytes()
	T.Log("no of bytes used for 64x64 Matrix is: ", bytes)
	if bytes > 8*64*64 {
		T.Error("this can not be true")
	}
}
