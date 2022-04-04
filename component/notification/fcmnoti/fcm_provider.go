package fcmnoti

import (
	"context"
	"github.com/NaySoftware/go-fcm"
	log "github.com/sirupsen/logrus"
)

type fcmProvider struct {
	token string
}

func NewFcmProvider(token string) *fcmProvider {
	return &fcmProvider{
		token: token,
	}
}

func (fcmP *fcmProvider) PushNotification(ctx context.Context, listToken []string, data interface{}) {
	dataNoti, ok := data.(*fcm.NotificationPayload)
	if !ok {
		log.Error("Cannot pressed type notification - ", data)
		return
	}
	f := fcm.NewFcmClient(fcmP.token)
	f.AppendDevices(listToken)
	f.SetNotificationPayload(dataNoti)

	_, err := f.Send()
	if err != nil {
		log.Error(err)
		return
	}
}
