package notification

import "fmt"

type Notifier interface {
	Send()
}
type SMS struct {
	PhoneNumber string
}
type Email struct {
	Address string
}

func (s SMS) Send() {
	fmt.Printf("SMS sent to: %s", s.PhoneNumber)
}
func (e Email) Send() {
	fmt.Printf("Email sent to: %s", e.Address)
}
func ExecuteSend(n Notifier) {
	n.Send()
}
