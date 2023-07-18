package vendors

import "fmt"

type Stripe struct {
	paymentHistory []string
}

func NewStripe() *Stripe {
	return &Stripe{}
}

func (stripe *Stripe) Pay(amount float32) string {
	paymentMessage := fmt.Sprintf("Paid %v amout via Stripe", amount)
	stripe.paymentHistory = append(stripe.paymentHistory, paymentMessage)
	return paymentMessage
}

func (stripe *Stripe) GetPaymentHistory() string {
	var paymentHistory string

	for _, p := range stripe.paymentHistory {
		paymentHistory += p + "\n"
	}

	return paymentHistory
}
