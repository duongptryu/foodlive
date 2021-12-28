package subscribe

import (
	"context"
	"fooddelivery/component"
)

func SetupSubscriber(appCtx component.AppContext) {
	ctx := context.Background()
	IncreaseLikeCountRestaurant(appCtx, ctx)
	DecreaseLikeCountRestaurant(appCtx, ctx)
}
