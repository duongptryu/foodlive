package subscribe

import (
	"context"
	"foodlive/component"
)

func SetupSubscriber(appCtx component.AppContext) {
	ctx := context.Background()
	IncreaseLikeCountRestaurant(appCtx, ctx)
	DecreaseLikeCountRestaurant(appCtx, ctx)
	CalculateRatingRestaurantWhenUserCreate(ctx, appCtx)
}
