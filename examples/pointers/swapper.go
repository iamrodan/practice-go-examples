package pointers

// Directly swap values without returning
// import it to main as swapper "practice-go-examples/examples/pointers" if this package name is different that folder name i.e package swapper
func SwapTwoInts(a *int, b *int) {
	temp := *b
	*b = *a
	*a = temp
}
