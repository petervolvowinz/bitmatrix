package bitmatrix

type Matrix interface {
	GetMatrix(n,m int) Matrix
	SetIndex(i, j int, val bool)
	GetIndex(i, j int) bool
	PrintMatrix()
	Multiply(B Matrix) Matrix
	Add(B Matrix) Matrix
	GetDimensions() (int,int)
	GetSizeInBytes() int
}
