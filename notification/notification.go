package notification

import "fmt"

type Notifier interface {
	Send() string
}
type SMS struct {
	PhoneNumber string
}
type Email struct {
	Address string
}

func (s SMS) Send() string {
	return "Sent SMS to: " + s.PhoneNumber
}
func (e Email) Send() string {
	return "Sent Email to: " + e.Address
}
func ExecuteNotification(n Notifier, s []string) {
	fmt.Println(n.Send())
}
