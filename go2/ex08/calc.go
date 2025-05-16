package ex08

func Sum(in ...int) int {
	sum := 0
	for i := 0; i < len(in); i++ {
		sum += in[i]
	}
	return sum
}
