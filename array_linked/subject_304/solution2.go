package subject_304

// NumMatrix2 304. 二维区域和检索 - 矩阵不可变
type NumMatrix2 struct {
	Matrix [][]int
}

func Constructor2(matrix [][]int) NumMatrix2 {
	m := len(matrix[0]) + 1
	sumArray := make([][]int, len(matrix)+1)
	for i := range sumArray {
		sumArray[i] = make([]int, m)
	}
	for i := 1; i < len(sumArray); i++ {
		for j := 1; j < m; j++ {
			sumArray[i][j] = sumArray[i-1][j] + sumArray[i][j-1] - sumArray[i-1][j-1] + matrix[i-1][j-1]
		}
	}
	return NumMatrix2{
		Matrix: sumArray,
	}
}

func (this *NumMatrix2) SumRegion2(row1 int, col1 int, row2 int, col2 int) int {
	return this.Matrix[row2+1][col2+1] - this.Matrix[row2+1][col1] - this.Matrix[row1][col2+1] + this.Matrix[row1][col1]
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2); 2 1 4 3
 */
