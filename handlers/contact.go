package handlers

import (
	"fmt"
	"net/http"
	"net/smtp"
	"os"

	"github.com/go-playground/validator"
	"github.com/svbrodersen/mySite/views/components"
	"github.com/svbrodersen/mySite/views/templates"
)

type ContactForm struct {
	Email   string `json:"Email" validate:"required,email"`
	Subject string `json:"Subject" validate:"required"`
	Body    string `json:"Body" validate:"required"`
}

func send_mail_to_self(form *ContactForm) error {
	from := "simon@3450.dk"
	password := string(os.Getenv("EMAIL_PASSWORD"))
	to := []string{from}

	message := []byte(fmt.Sprintf(
		"To: simon@3450.dk\r\n"+
			"From: %s\r\n"+
			"Subject: %s\r\n"+
			"%s\r\n", form.Email, form.Subject, form.Body,
	))
	auth := smtp.PlainAuth("", from, password, "smtp.gmail.com")

	err := smtp.SendMail("smtp.gmail.com:587", auth, from, to, message)
	return err
}

func ContactHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PATCH" {
		templates.Contact(true).Render(r.Context(), w)
	} else if r.Method == "GET" {
		templates.Contact(false).Render(r.Context(), w)
	} else if r.Method == "POST" {
		contact_form := ContactForm{
			Email:   r.FormValue("email"),
			Subject: r.FormValue("subject"),
			Body:    r.FormValue("body"),
		}

		validate := validator.New()
		err := validate.Struct(contact_form)
		if err != nil {
			fmt.Printf(err.Error())
			components.ContactComponent(false).Render(r.Context(), w)
			return
		}
		err = send_mail_to_self(&contact_form)
		if err != nil {
			fmt.Printf(err.Error())
			components.ContactComponent(false).Render(r.Context(), w)
			return
		}
		components.ContactComponent(true).Render(r.Context(), w)
	}
}
