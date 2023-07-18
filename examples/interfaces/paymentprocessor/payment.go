package paymentprocessor

import (
	"practice-go-examples/examples/interfaces/paymentprocessor/vendors"
)

type PaymentProcessor interface {
	Pay(amount float32) string
	GetPaymentHistory() string
}

func GetPaymentMethod(name string) PaymentProcessor {
	var paymentMethod PaymentProcessor
	switch name {
	case "stripe":
		paymentMethod = vendors.NewStripe()
	case "paypal":
		paymentMethod = vendors.NewPaypal()
	}

	return paymentMethod

}

func PayAmount(amount float32, paymentMethod PaymentProcessor) string {
	result := paymentMethod.Pay(amount)
	return result
}

func GetPaymentHistory(paymentMethod PaymentProcessor) string {
	result := paymentMethod.GetPaymentHistory()
	return result
}
