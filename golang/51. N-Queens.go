package golang

import "strings"

type checker struct {
	n            int   // n-queens
	rows         []int // 用于标记是否被当前所在行的某一列方向的皇后被攻击
	mainDiagonal []int // 用于标记是否被主对角线方向的皇后攻击
	antiDiagonal []int // 用于标记是否被次对角线方向的皇后攻击
	queens       []int // (i,queens[i]) 表示将皇后放置在 i 行，queens[i] 列
	solutions    [][]string
}

func newChecker(n int) *checker {
	return &checker{
		n:            n,
		rows:         make([]int, n),
		mainDiagonal: make([]int, 2*n-1),
		antiDiagonal: make([]int, 2*n-1),
		queens:       make([]int, n),
	}
}

func (c *checker) backtrack(row int) {
	if row < c.n {
		for col := 0; col < c.n; col++ {
			if c.queenIsNotUnderAttackAt(row, col) {
				c.placeQueen(row, col)
				if row == c.n-1 {
					c.addSolution()
				}
				c.backtrack(row + 1)
				c.removeQueen(row, col)
			}
		}
	}
}

func (c checker) queenIsNotUnderAttackAt(row, col int) bool {
	res := c.rows[col] + c.mainDiagonal[row-col+c.n-1] + c.antiDiagonal[row+col]
	return res == 0
}

func (c *checker) placeQueen(row, col int) {
	c.queens[row] = col
	c.rows[col] = 1
	c.mainDiagonal[row-col+c.n-1] = 1
	c.antiDiagonal[row+col] = 1
}

func (c *checker) removeQueen(row, col int) {
	c.queens[row] = 0
	c.rows[col] = 0
	c.mainDiagonal[row-col+c.n-1] = 0
	c.antiDiagonal[row+col] = 0
}

func (c *checker) addSolution() {
	solution := make([]string, 0, c.n)
	rb := &strings.Builder{}
	for _, col := range c.queens {
		for i := 0; i < col; i++ {
			rb.WriteByte('.')
		}
		rb.WriteByte('Q')
		for i := col + 1; i < c.n; i++ {
			rb.WriteByte('.')
		}
		solution = append(solution, rb.String())
		rb.Reset()
	}
	c.solutions = append(c.solutions, solution)
}

func (c *checker) getSolutions() [][]string {
	return c.solutions
}

func solveNQueens(n int) [][]string {
	c := newChecker(n)
	c.backtrack(0)
	return c.getSolutions()
}
