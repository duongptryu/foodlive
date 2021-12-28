package subscribe

import (
	"context"
	"fooddelivery/common"
	"fooddelivery/component"
	"fooddelivery/modules/restaurant/restaurantstore"
	log "github.com/sirupsen/logrus"
)

func DecreaseLikeCountRestaurant(appCtx component.AppContext, ctx context.Context) {
	c, _ := appCtx.GetPubSubProvider().Subscribe(ctx, common.TopicUserDisLikeRestaurant)

	restaurantStore := restaurantstore.NewSqlStore(appCtx.GetDatabase())

	go func() {
		defer common.AppRecovery()

		msg := <-c

		likeData := msg.Data().(CastingRestaurant)

		err := restaurantStore.DecreaseLikeCount(ctx, likeData.GetRestaurantId())
		if err != nil {
			log.Error(err)
		}
	}()
}
