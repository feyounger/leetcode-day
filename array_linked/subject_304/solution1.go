package subject_304

type NumMatrix struct {
	Matrix [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	return NumMatrix{
		Matrix: matrix,
	}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	var sum int
	for i := row1; i < row2+1; i++ {
		for j := col1; j < col2+1; j++ {
			sum += this.Matrix[i][j]
		}
	}
	return sum
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2); 2 1 4 3
 */
