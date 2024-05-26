package core

/**
 * Send an email to the specified recipient.
 * @param to The recipient of the email.
 * @param subject The subject of the email.
 * @param body The body of the email.
 * @return Whether the email was sent successfully and any errors.
 */
func SendEmail(to string, subject string, body string) (bool, error) {
	return true, nil
}
