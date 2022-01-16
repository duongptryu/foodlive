package subscribe

import (
	"context"
	"foodlive/common"
	"foodlive/component"
	"foodlive/modules/restaurant/restaurantstore"
	log "github.com/sirupsen/logrus"
)

type CastingRestaurant interface {
	GetRestaurantId() int
	GetOwnerId() int
}

func IncreaseLikeCountRestaurant(appCtx component.AppContext, ctx context.Context) {
	c, _ := appCtx.GetPubSubProvider().Subscribe(ctx, common.TopicUserLikeRestaurant)

	restaurantStore := restaurantstore.NewSqlStore(appCtx.GetDatabase())

	go func() {
		defer common.AppRecovery()

		msg := <-c

		likeData := msg.Data().(CastingRestaurant)

		err := restaurantStore.IncreaseLikeCount(ctx, likeData.GetRestaurantId())
		if err != nil {
			log.Error(err)
		}
	}()
}
