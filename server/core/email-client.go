package core

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"
)

/**
 * Send an email to the specified recipient.
 * @param to The recipient of the email.
 * @param subject The subject of the email.
 * @param body The body of the email.
 * @return Whether the email was sent successfully and any errors.
 */
func SendEmail(to string, subject string, body string) (bool, error) {
	token := "<secret_token>"
	httpHost := "https://send.api.mailtrap.io/api/send"
	message := []byte(fmt.Sprintf(`{"from":{"email":"notiflash@zentools.in"},

"to":[{"email":"%s"}],

"subject":"NotiFlash ⚡️",

"text":"%s"}`, to, body))

	request, err := http.NewRequest(http.MethodPost, httpHost, bytes.NewBuffer(message))

	if err != nil {
		fmt.Println("Error creating request", err)
		return false, err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+token)

	client := http.Client{Timeout: 5 * time.Second}
	res, err := client.Do(request)

	if err != nil {
		fmt.Println("Error sending request", err)
		return false, err
	}

	defer res.Body.Close()

	responseBody, err := io.ReadAll(res.Body)
	if err != nil {

		fmt.Println("Failed to read email response body.", err)
		return false, err
	}

	fmt.Println(string(responseBody))
	return true, nil
}
