package mail

import (
	"testing"

	"github.com/maliByatzes/parte-server/config"
	"github.com/stretchr/testify/require"
)

func TestSendEmailWithGmail(t *testing.T) {
	config, err := config.LoadConfig()
	require.NoError(t, err)

	sender := NewGmailSender(config.Mail.Sender, config.Mail.SenderAddress, config.Mail.Password)

	subject := "A test email"
	content := `
	<h1>Hello User!</h1>
	<p>This is a test message from <a href="https://youtu.be/dQw4w9WgXcQ?si=CJ-DYhCnV6ARuysB">Parte</a>.</p>
	`
	to := []string{"malib2027@gmail.com"}
	attachFiles := []string{"../sqlc.yaml"}

	err = sender.SendEmail(subject, content, to, nil, nil, attachFiles)
	require.NoError(t, err)
}
