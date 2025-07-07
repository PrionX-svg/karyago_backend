package pkg

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/smtp"
	"os"
)

func SendEmail(to, subject, body string) error {
	host := os.Getenv("MAIL_HOST")
	port := os.Getenv("MAIL_PORT") // as string, e.g., "587"
	username := os.Getenv("MAIL_USERNAME")
	password := os.Getenv("MAIL_PASSWORD")
	fromName := os.Getenv("MAIL_NAME")
	useSSL := os.Getenv("MAIL_SSL") == "true"
	useTLS := os.Getenv("MAIL_TLS") == "true"

	addr := fmt.Sprintf("%s:%s", host, port)

	from := username
	if fromName == "" {
		fromName = username
	}

	fromHeader := fmt.Sprintf("From: \"%s\" <%s>\r\n", fromName, from)
	toHeader := fmt.Sprintf("To: %s\r\n", to)
	subjectHeader := fmt.Sprintf("Subject: %s\r\n", subject)

	message := []byte(
		fromHeader +
			toHeader +
			subjectHeader +
			"MIME-Version: 1.0\r\n" +
			"Content-Type: text/html; charset=\"UTF-8\"\r\n\r\n" +
			body,
	)

	auth := smtp.PlainAuth("", username, password, host)

	if useSSL {
		tlsConfig := &tls.Config{
			InsecureSkipVerify: true,
			ServerName:         host,
		}
		conn, err := tls.Dial("tcp", addr, tlsConfig)
		if err != nil {
			return fmt.Errorf("SSL connection error: %v", err)
		}
		client, err := smtp.NewClient(conn, host)
		if err != nil {
			return fmt.Errorf("SMTP client error (SSL): %v", err)
		}
		defer client.Quit()

		if err := client.Auth(auth); err != nil {
			return fmt.Errorf("auth failed: %v", err)
		}
		if err := client.Mail(from); err != nil {
			return err
		}
		if err := client.Rcpt(to); err != nil {
			return err
		}
		w, err := client.Data()
		if err != nil {
			return err
		}
		defer w.Close()

		_, err = w.Write(message)
		return err
	}

	if useTLS {
		// STARTTLS
		return smtp.SendMail(addr, auth, from, []string{to}, message)
	}

	// Plain (not recommended)
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return fmt.Errorf("plain connection error: %v", err)
	}
	client, err := smtp.NewClient(conn, host)
	if err != nil {
		return fmt.Errorf("SMTP client error: %v", err)
	}
	defer client.Quit()

	if err := client.Auth(auth); err != nil {
		return fmt.Errorf("auth failed: %v", err)
	}
	if err := client.Mail(from); err != nil {
		return err
	}
	if err := client.Rcpt(to); err != nil {
		return err
	}
	w, err := client.Data()
	if err != nil {
		return err
	}
	defer w.Close()

	_, err = w.Write(message)
	return err
}
