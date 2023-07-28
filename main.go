package main

import (
	"fmt"
	"practice-go-examples/examples/concurrency/channels"
	"practice-go-examples/examples/embedding/guards"
	"practice-go-examples/examples/embedding/logger"
	"practice-go-examples/examples/interfaces/paymentprocessor"
	"practice-go-examples/examples/panic"
)

func main() {
	// SitesAvailabilityCheckExample()
	// paymentProcessorExampleWithInterfaces()
	// embeddedLoggerExample()
	runGoPanicExample()
}

func runGoPanicExample() {
	result_channel := make(chan int)
	msg_channel := make(chan string)
	go panic.Sum(5, -10, result_channel)
	go panic.SayHelloMsg(msg_channel)
	fmt.Printf("sayHelloMsg output: %v", <-msg_channel)
	fmt.Printf("Sum of %v + %v is %v\n", 5, 10, <-result_channel)
}

func SitesAvailabilityCheckExample() {
	urls := []string{
		"https://www.easyjet.com/",
		"https://www.skyscanner.de/",
		"https://www.ryanair.com",
		"https://wizzair.com/",
		"https://www.swiss.com/",
	}
	channels.SitesAvailabilityCheck(urls)
}

func embeddedLoggerExample() {
	user1 := guards.NewUser(1, true)
	user2 := guards.NewUser(2, false)
	guard1 := guards.NewGuard(user1)
	guard2 := guards.NewGuard(user2)
	logger1 := logger.NewLogger(guard1)
	logger2 := logger.NewLogger(guard2)
	logger1.LogInfo("Logger 1 ko log")
	logger2.LogInfo("Logger 2 ko log") // should throw no permission

}

func paymentProcessorExampleWithInterfaces() {
	// Pay 10, 20, 30 to stripe
	const (
		stripe = "stripe"
		paypal = "paypal"
	)

	stripePaymentProcessor := paymentprocessor.GetPaymentMethod(stripe)
	paypalPaymentProcessor := paymentprocessor.GetPaymentMethod(paypal)
	for _, amount := range []int{10, 20, 30} {
		fmt.Println(paymentprocessor.PayAmount(float32(amount), stripePaymentProcessor))
	}
	fmt.Printf("\n\nPayment History:\n")
	fmt.Println(paymentprocessor.GetPaymentHistory(stripePaymentProcessor))

	// pay 5, 15, 25 to paypal
	for _, amount := range []int{5, 15, 25} {
		fmt.Println(paymentprocessor.PayAmount(float32(amount), paypalPaymentProcessor))
	}
	fmt.Printf("\n\nPayment History:\n")
	fmt.Println(paymentprocessor.GetPaymentHistory(paypalPaymentProcessor))
}
