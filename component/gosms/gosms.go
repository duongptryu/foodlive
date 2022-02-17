package gosms

import (
	"context"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

type GoSMS interface {
	SendOTP(ctx context.Context, phoneNumber string, otp string) error
}

type goSms struct {
	client        *twilio.RestClient
	myPhoneNumber string
}

func NewGoSms(accountSid string, authToken string, myPhoneNumber string) *goSms {
	if accountSid == "" || authToken == "" || myPhoneNumber == "" {
		log.Fatal("accountSid or authToken or phone number is empty")
	}
	client := twilio.NewRestClientWithParams(twilio.RestClientParams{
		Username: accountSid,
		Password: authToken,
	})
	return &goSms{
		client,
		myPhoneNumber,
	}
}

func (c *goSms) SendOTP(ctx context.Context, phoneNumber string, otp string) error {
	params := &openapi.CreateMessageParams{}
	params.SetTo("+" + phoneNumber)
	params.SetFrom(c.myPhoneNumber)
	params.SetBody("Welcome to FoodHub. Your OTP is: " + otp + " .This OTP is exist in 60s")

	resp, err := c.client.ApiV2010.CreateMessage(params)
	if err != nil {
		fmt.Println(err.Error())
		return err
	} else {
		response, _ := json.Marshal(*resp)
		fmt.Println("Response: " + string(response))
	}
	return nil
}
