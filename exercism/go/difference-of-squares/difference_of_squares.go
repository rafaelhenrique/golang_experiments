package diffsquares

// Reference about calculation https://brilliant.org/wiki/sum-of-n-n2-or-n3/

// SquareOfSum calculates square of the sum of the first N natural numbers
func SquareOfSum(n int) (squareOfSum int) {
	squareOfSum = (n * (n + 1)) / 2
	return squareOfSum * squareOfSum
}

// SumOfSquares calculates sum of the squares of the first N natural numbers
func SumOfSquares(n int) (sumOfSquares int) {
	sumOfSquares = (n * (n + 1) * (2*n + 1)) / 6
	return sumOfSquares
}

// Difference calculate difference of the squares of a number N
func Difference(n int) (difference int) {
	return SquareOfSum(n) - SumOfSquares(n)
}
