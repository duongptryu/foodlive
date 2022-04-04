package notification

import "context"

type NotiProvider interface {
	PushNotification(ctx context.Context, listToken []string, data interface{})
}
