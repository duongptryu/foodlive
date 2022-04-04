package subscribe

import (
	"context"
	"foodlive/common"
	"foodlive/component"
	"foodlive/modules/restaurant/restaurantstore"
	log "github.com/sirupsen/logrus"
)

func DecreaseLikeCountRestaurant(appCtx component.AppContext, ctx context.Context) {
	c, _ := appCtx.GetPubSubProvider().Subscribe(ctx, common.TopicUserDisLikeRestaurant)

	restaurantStore := restaurantstore.NewSqlStore(appCtx.GetDatabase())

	go func() {
		defer common.AppRecovery()
		for {
			msg := <-c

			likeData := msg.Data().(CastingRestaurant)

			err := restaurantStore.DecreaseLikeCount(ctx, likeData.GetRestaurantId())
			if err != nil {
				log.Fatalln(err)
			}
		}
	}()
}
