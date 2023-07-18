package examples

// takes multiple args of nums and return their sum
func VariadicFunctionAdder(nums ...int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	return total
}
