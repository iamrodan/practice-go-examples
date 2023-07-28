package panic

import (
	"fmt"
	"time"
)

func Sum(a, b int, result_channel chan int) {
	time.Sleep(2 * time.Second)
	defer func() {
		fmt.Println("Defer 1 function")
	}()
	defer func() {
		fmt.Println("Defer 2 Func")
	}()
	if a < 0 || b < 0 {
		panic("Got non-positive number in the input")
	}
	defer func() {
		fmt.Println("Defer after if")
	}()
	sum := a + b
	result_channel <- sum
}

func SayHelloMsg(msg_channel chan string) {
	time.Sleep(3 * time.Second)
	msg_channel <- "Hello from sayHelloMsg()"
}
