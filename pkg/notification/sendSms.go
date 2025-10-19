package notification

import (
	"encoding/json"
	"errors"
	"fmt"

	"go-ecommerce-app2/config"

	"github.com/twilio/twilio-go"

	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

// port
type NotificationClient interface {
	SendSMS(phone string, message string) error
}

// repo
type notificationClient struct {
	config config.AppConfig
}

// Twilio
func (nc notificationClient) SendSMS(phone string, message string) error {
	accountSid := nc.config.TwilioAccountSID
	authToken := nc.config.TwilioAuthToken
	phoneFrom := nc.config.TwilioFromPhoneNumber

	println("phone number is passed successfully")
	println(phoneFrom)
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

	if phone == "" {
		return errors.New("empty phone number not allowed")
	}

	params := &twilioApi.CreateMessageParams{}
	// carefull only pass verified phone number only for free trial
	// also remember to pass country code in phone number
	params.SetTo(phone)
	params.SetFrom("+13613013165")
	params.SetBody(message)

	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println("Error sending SMS message: " + err.Error())
		return errors.New("error sending SMS")
	} else {
		response, _ := json.Marshal(*resp)
		fmt.Println("Response: " + string(response))
	}
	return nil
}

// constructor
func NewNotificationClient(config config.AppConfig) NotificationClient {
	return &notificationClient{
		config,
	}
}
