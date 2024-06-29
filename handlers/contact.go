package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/smtp"
	"os"

	"github.com/go-playground/validator"
	"github.com/svbrodersen/mySite/views/components"
	"github.com/svbrodersen/mySite/views/templates"
)

type ContactForm struct {
	Name    string `json:"Name" validate:"required"`
	Email   string `json:"Email" validate:"required, email"`
	Subject string `json:"Subject" validate:"required"`
	Body    string `json:"Body" validate:"required"`
}

func send_mail_to_self(form *ContactForm) error {
	from := "simon@3450.dk"
	password := os.Getenv("EMAIL_PASSWORD")

	to := []string{from}

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	message := []byte("This is a test email")
	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(fmt.Sprintf("%s:%s", smtpHost, smtpPort), auth, from, to, message)
	return err
}

func ContactHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PATCH" {
		templates.Contact(true).Render(r.Context(), w)
	} else if r.Method == "GET" {
		templates.Contact(false).Render(r.Context(), w)
	} else if r.Method == "POST" {
		var contact_form ContactForm
		err := json.NewDecoder(r.Body).Decode(&contact_form)
		if err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}
		validate := validator.New()
		err = validate.Struct(contact_form)
		if err != nil {
			errors := err.(validator.ValidationErrors)
			http.Error(w, fmt.Sprintf("Validation error: %s", errors), http.StatusBadRequest)
			return
		}
		err = send_mail_to_self(&contact_form)
		if err != nil {
			http.Error(w, "Unable to send email to myself", http.StatusBadRequest)
		}
		components.ContactSuccess().Render(r.Context(), w)
	}
}
