package main

func main() {
	product := "car"
	payWay := 3

	var payment Payment
	switch payWay {
	case 1:
		payment = NewCardPayment("12345", "111")
	case 2:
		payment = NewPayPalPayment()
	case 3:
		payment = NewQIWIPayment()
	}

	processOrder(product, payment)
}

func processOrder(product string, payment Payment) {
	err := payment.Pay()
	if err != nil {
		return
	}

}

type Payment interface {
	Pay() error
}

type CardPayment struct {
	cardNumber string
	cvv string
}

func NewCardPayment(cardNumber string, cvv string) Payment{
	return &CardPayment{
		cardNumber: cardNumber,
		cvv: cvv,
	}
}

func NewPayPalPayment() Payment{
	return &PayPalPayment{}
}

func NewQIWIPayment() Payment{
	return &QIWIPayment{}
}

type PayPalPayment struct {
}

type QIWIPayment struct {
}

func (p *CardPayment) Pay() error  {
	println("Payed by Card: ", p.cardNumber, p.cvv)
	return nil
}

func (p *PayPalPayment) Pay() error  {
	println("Payed by PayPal")
	return nil
}

func (p *QIWIPayment) Pay() error  {
	println("Payed by QIWI")
	return nil
}