package diffsquares

// SquareOfSum calculates square of the sum of the first N natural numbers
func SquareOfSum(n int) (squareOfSum int) {
	for i := 1; i <= n; i++ {
		squareOfSum += i
	}
	return squareOfSum * squareOfSum
}

// SumOfSquares calculates sum of the squares of the first N natural numbers
func SumOfSquares(n int) (sumOfSquares int) {
	for i := 1; i <= n; i++ {
		sumOfSquares += i * i
	}
	return sumOfSquares
}

// Difference calculate difference of the squares of a number N
func Difference(n int) (difference int) {
	return SquareOfSum(n) - SumOfSquares(n)
}
