package solution

func SwapVariable1(a *int, b *int) {
	*a, *b = *b, *a
}

func SwapVariable2(a *int, b *int) {
	*a = *a + *b
	*b = *a - *b
	*a = *a - *b
}
