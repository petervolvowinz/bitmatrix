package bitmatrix

type Matrix interface {
	GetMatrix(size int) Matrix
	SetIndex(i, j int, val bool)
	GetIndex(i, j int) bool
	PrintMatrix()
	Multiply(B Matrix) Matrix
	Add(B Matrix) Matrix
	GetDimensions() int
	GetSizeInBytes() int
}
