package email

import (
	"context"
	"log/slog"
)

// Sender defines the interface for sending emails.
type Sender interface {
	Send(ctx context.Context, to, subject, body string) error
}

// StubSender is a stub implementation that logs emails to stdout.
type StubSender struct {
	logger *slog.Logger
}

// NewStubSender creates a new stub email sender.
func NewStubSender(logger *slog.Logger) *StubSender {
	return &StubSender{
		logger: logger,
	}
}

// Send logs the email to stdout without actually sending it.
func (s *StubSender) Send(ctx context.Context, to, subject, body string) error {
	s.logger.InfoContext(ctx, "stub email sent",
		"to", to,
		"subject", subject,
		"body", body,
	)
	return nil
}
