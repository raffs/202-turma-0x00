package extra

func Fatorial(n int) int {
	if n <= 1 {
		return 1
	}
	return Fatorial(n - 1) * n
}
