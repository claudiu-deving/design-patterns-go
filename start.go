package start

import "fmt"

func main() {
	fmt.Print("1. Hello World\n")
	implicitVariableInt := 10
	fmt.Printf("2. I am a number %d\n", implicitVariableInt)
	implicitstring := "Am I a string?"
	fmt.Printf("3. %s\n", implicitstring)

	result1, result2 := multipleReturns(1, 4, 5)

	fmt.Printf("4.Multiple returns example:\n%d %d", result1, result2)
}

func multipleReturns(a, b, c int) (int, int) {
	return a + b, a + c
}
