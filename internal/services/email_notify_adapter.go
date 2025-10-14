package services

import "hris_backend/pkg"

type SMTPNotifier struct{}

func NewSMTPNotifier() NotificationService { return &SMTPNotifier{} }

func (n *SMTPNotifier) SendEmail(to, subject, htmlBody string) error {
	return pkg.SendEmail(to, subject, htmlBody)
}
