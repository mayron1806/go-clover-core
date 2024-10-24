package service

import (
	"fmt"
	"net/smtp"

	"github.com/mayron1806/go-clover-core/config"
	"github.com/mayron1806/go-clover-core/logger"
)

type EmailOptions struct {
	ShouldSendEmails bool   `env:"SHOULD_SEND_EMAILS" default:"false"`
	SmtpHost         string `env:"SMTP_HOST" default:"localhost"`
	SmtpPort         string `env:"SMTP_PORT" default:"25"`
	SmtpUser         string `env:"SMTP_USER"`
	SmtpPass         string `env:"SMTP_PASS"`
	SmtpFrom         string `env:"SMTP_FROM"`
}
type EmailService struct {
	auth      smtp.Auth
	sendEmail bool
	address   string
	from      string
	logger    *logger.Logger
}

func NewEmailService() (*EmailService, error) {
	logger := logger.NewLogger(logger.LoggerOptions{Prefix: "EMAIL"})
	envLoader := config.NewEnvLoader[EmailOptions]()
	env, err := envLoader.LoadEnv()
	if err != nil {
		logger.Errorf("error loading environment variables: %s", err.Error())
		return nil, err
	}
	if !env.ShouldSendEmails {
		return &EmailService{
			logger:    logger,
			sendEmail: false,
		}, nil
	}

	// Verificar se as variáveis de ambiente estão configuradas corretamente
	if env.SmtpHost == "" || env.SmtpPort == "" {
		logger.Errorf("SMTP_HOST or SMTP_PORT is not set")
		return nil, err
	}

	auth := smtp.PlainAuth("", env.SmtpUser, env.SmtpPass, env.SmtpHost)
	address := fmt.Sprintf("%s:%s", env.SmtpHost, env.SmtpPass)

	return &EmailService{
		logger:    logger,
		auth:      auth,
		sendEmail: true,
		address:   address,
		from:      env.SmtpFrom,
	}, nil
}
func (e *EmailService) SendEmail(to string, subject string, body string) error {
	if !e.sendEmail {
		e.logger.Debug("email sending is disabled")
		e.logger.Debugf("to: %s, subject: %s, body: %s", to, subject, body)
		return nil
	}
	err := smtp.SendMail(e.address, e.auth, e.from, []string{to}, []byte(body))
	if err != nil {
		e.logger.Errorf("error sending email: %s", err.Error())
		return err
	}
	return nil
}
