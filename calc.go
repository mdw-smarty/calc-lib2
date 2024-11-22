package calc_lib2

type Addition struct{}

func (this *Addition) Calculate(a, b int) int {
	return a + b
}
