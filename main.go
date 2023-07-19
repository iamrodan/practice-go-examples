package main

import (
	"fmt"
	"practice-go-examples/examples/concurrency/waitgroups"
	"practice-go-examples/examples/embedding/guards"
	"practice-go-examples/examples/embedding/logger"
	"practice-go-examples/examples/interfaces/paymentprocessor"
)

func main() {
	SitesAvailabilityCheckExample()
	// paymentProcessorExampleWithInterfaces()
	// embeddedLoggerExample()
}

func SitesAvailabilityCheckExample() {
	urls := []string{
		"https://www.easyjet.com/",
		"https://www.skyscanner.de/",
		"https://www.ryanair.com",
		"https://wizzair.com/",
		"https://www.swiss.com/",
	}
	waitgroups.SitesAvailabilityCheck(urls)
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
