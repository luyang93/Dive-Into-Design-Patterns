package main

import "fmt"

type Payment struct {
	amount float64
}

func (p *Payment) getAmount() float64 {
	return p.amount
}

type RefundStatus int

const (
	SUCCESS RefundStatus = 1
	FAILURE RefundStatus = 2
)

type Billing struct{}

func (b *Billing) GetPaymentForOrder(orderID int) Payment {
	return Payment{amount: 50.0}
}

func (b *Billing) ProcessRefund(p Payment) RefundStatus {
	if p.getAmount() > 0 {
		return SUCCESS
	} else {
		return FAILURE
	}
}

type Shipping struct{}

func (s *Shipping) UpdateShippingAddress(orderID int, newAddress string) {
	fmt.Printf("Shipping address for order %d updated to: %s\n", orderID, newAddress)
}

type Issue struct {
	Description string
}

func (i *Issue) GetDescription() string {
	return i.Description
}

type CustomerService struct{}

func (c CustomerService) NotifyCustomer(message string) {
	fmt.Printf("Notification sent to customer: %s\n", message)
}

func (c CustomerService) EscalateToManager(issue Issue) {
	fmt.Printf("Issue escalated to manager: %s\n", issue.GetDescription())
}

type CustomerSupportFacade struct {
	Billing         Billing
	Shipping        Shipping
	CustomerService CustomerService
}

func (csf *CustomerSupportFacade) HandleRefundRequest(orderID int) {
	payment := csf.Billing.GetPaymentForOrder(orderID)
	refundStatus := csf.Billing.ProcessRefund(payment)
	csf.CustomerService.NotifyCustomer(fmt.Sprintf("Refund status: %d", refundStatus))
}

func (csf *CustomerSupportFacade) ChangeShippingAddress(orderID int, newAddress string) {
	csf.Shipping.UpdateShippingAddress(orderID, newAddress)
	csf.CustomerService.NotifyCustomer("Shipping address updated")
}

func (csf *CustomerSupportFacade) EscalateToManager(issue Issue) {
	csf.CustomerService.EscalateToManager(issue)
}

func main() {
	billing := Billing{}
	shipping := Shipping{}
	customerService := CustomerService{}

	customerSupport := CustomerSupportFacade{
		Billing:         billing,
		Shipping:        shipping,
		CustomerService: customerService,
	}

	customerSupport.HandleRefundRequest(12345)
	customerSupport.ChangeShippingAddress(12345, "123 New Street, New York, NY 10001")
	issue := Issue{Description: "Product not working properly"}
	customerSupport.EscalateToManager(issue)
}
