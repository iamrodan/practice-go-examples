package vendors

import "fmt"

type Paypal struct {
	paymentHistory []string
}

func NewPaypal() *Paypal {
	return &Paypal{}
}

func (paypal *Paypal) Pay(amount float32) string {
	paymentMessage := fmt.Sprintf("Paid %v amout via Paypal", amount)
	paypal.paymentHistory = append(paypal.paymentHistory, paymentMessage)
	return paymentMessage
}

func (paypal *Paypal) GetPaymentHistory() string {
	var paymentHistory string

	for _, p := range paypal.paymentHistory {
		paymentHistory += p + "\n"
	}

	return paymentHistory
}
