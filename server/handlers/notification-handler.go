package handlers

import (
	"encoding/json"
	"main/core"
	"net/http"
)

type NotificationRequest struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

func NotificationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "only POST is allowed", http.StatusMethodNotAllowed)
		return
	}

	var request NotificationRequest
	body := json.NewDecoder(r.Body).Decode(&request)

	if body != nil {
		http.Error(w, body.Error(), http.StatusBadRequest)
		return
	}

	if request.To == "" || request.Subject == "" || request.Body == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	// Send the email
	success, err := core.SendEmail(request.To, request.Subject, request.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if success {
		w.WriteHeader(http.StatusOK)
		return
	} else {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
