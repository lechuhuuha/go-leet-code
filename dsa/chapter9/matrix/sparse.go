package matrix

type Cell struct {
	Row    int
	Column int
	Value  float64
}

type SparseMatrix struct {
	Cells []Cell
	shape [2]int
}

func (s *SparseMatrix) Shape() (int, int) {
	return s.shape[0], s.shape[1]
}

func (s *SparseMatrix) NumNonZero() int {
	return len(s.Cells)
}

func LessThan(cell Cell, i, j int) bool {
	if cell.Row < i && cell.Column < j {
		return true
	}
	return false
}

func Equal(cell Cell, i, j int) bool {
	if cell.Row == i && cell.Column == j {
		return true
	}
	return false
}

func (sparseMatrix *SparseMatrix) GetValue(i int, j int) float64 {
	var cell Cell
	for _, cell = range sparseMatrix.Cells {
		if LessThan(cell, i, j) {
			continue
		}
		if Equal(cell, i, j) {
			return cell.Value
		}
		return 0.0
	}
	return 0.0
}

func (sparseMatrix *SparseMatrix) SetValue(i int, j int, value float64) {
	var cell Cell
	var index int
	for index, cell = range sparseMatrix.Cells {
		if LessThan(cell, i, j) {
			continue
		}
		if Equal(cell, i, j) {
			sparseMatrix.Cells[index].Value = value
			return
		}
		sparseMatrix.Cells = append(sparseMatrix.Cells, Cell{})
		var k int
		for k = len(sparseMatrix.Cells) - 2; k >= index; k-- {
			sparseMatrix.Cells[k+1] = sparseMatrix.Cells[k]
		}
		sparseMatrix.Cells[index] = Cell{
			Row:    i,
			Column: j,
			Value:  value,
		}
		return
	}
	sparseMatrix.Cells = append(sparseMatrix.Cells, Cell{
		Row:    i,
		Column: j,
		Value:  value,
	})
}

func NewSparseMatrix(m int, n int) *SparseMatrix {
	return &SparseMatrix{
		Cells: []Cell{},
		shape: [2]int{m, n},
	}
}
