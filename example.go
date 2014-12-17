package main

import (
	"fmt"
	"github.com/collinglass/emailer/email"
)

// Generic Recipient
// Feel free to add any fields you'd like!
type Recipient struct {
	Name  string
	Email string
}

func main() {
	// Setup Emailer with your credentials
	emailer := email.Setup("GMAIL_EMAIL_ADDRESS", "GMAIL_PASSWORD")

	// Create email list of recipients
	emailList := []Recipient{
		Recipient{
			Name:  "Collin",
			Email: "collinglass@gmx.com",
		},
	}
	for _, rec := range emailList {

		// Write your message
		msg := fmt.Sprintf("Hey %s, \n\nI'm Collin, nice to meet you!\n\n--\nCheers,\n\nCollin Glass\n", rec.Name)

		// Send mailout
		err := emailer.SendMail(rec.Email, "I just sent this with Emailer", msg)

		// Handle Error
		if err != nil {
			fmt.Errorf("Error emailing %s: %s \n", rec.Email, err)
		} else {
			fmt.Printf("Successfully emailed %s with email %s \n", rec.Name, rec.Email)
		}
	}

}
