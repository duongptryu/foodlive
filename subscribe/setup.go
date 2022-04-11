package subscribe

import (
	"context"
	"foodlive/component"
)

func SetupSubscriber(appCtx component.AppContext) {
	ctx := context.Background()
	IncreaseLikeCountRestaurant(appCtx, ctx)
	DecreaseLikeCountRestaurant(appCtx, ctx)
	CalculateRatingRestaurant(ctx, appCtx)
	CalculateUpdateRatingRestaurant(ctx, appCtx)

	UserLikeFood(appCtx, ctx)
	DecreaseLikeCountFood(appCtx, ctx)
	CalculateRatingFood(ctx, appCtx)
}
