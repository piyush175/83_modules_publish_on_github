package calc

func Add(a ...int) int {
	sum := 0
	for i := range a {
		sum += i
	}
	return sum
}

func Sub(a ...int) int {
	sum := 0
	for i := range a {
		sum += i
	}
	return sum
}
