package subject_304

// NumMatrix2 304. 二维区域和检索 - 矩阵不可变
type NumMatrix2 struct {
	Matrix [][]int
}

func Constructor2(matrix [][]int) NumMatrix2 {
	sumArray := make([][]int, len(matrix)+1)
	sumArray[0] = make([]int, len(matrix[0])+1)
	for i := 1; i < len(matrix)+1; i++ {
		sumArray[i] = make([]int, len(matrix[i-1])+1)
		for j := 1; j < len(matrix[i-1])+1; j++ {
			sumArray[i-1][j] = sumArray[i-1][j-1] + matrix[i-1][j-1]
			if i > 1 {
				sumArray[i-1][j] += sumArray[i-2][j] - sumArray[i-2][j-1]
			}
		}
	}
	return NumMatrix2{
		Matrix: sumArray,
	}
}

func (this *NumMatrix2) SumRegion2(row1 int, col1 int, row2 int, col2 int) int {
	var sum int
	sum = this.Matrix[row2][col2+1] - this.Matrix[row2][col1]
	if row1 >= 1 {
		sum += this.Matrix[row1-1][col1] - this.Matrix[row1-1][col2+1]
	}
	return sum
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2); 2 1 4 3
 */
