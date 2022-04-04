package subscribe

import (
	"context"
	"foodlive/common"
	"foodlive/component"
	"foodlive/modules/restaurant/restaurantstore"
	"foodlive/modules/restaurantrating/restaurantratingstore"
	log "github.com/sirupsen/logrus"
)

type RatingData interface {
	GetRestaurantId() int
}

func CalculateRatingRestaurant(ctx context.Context, appCtx component.AppContext) {
	c, _ := appCtx.GetPubSubProvider().Subscribe(ctx, common.TopicUserCreateRestaurantRating)

	restaurantStore := restaurantstore.NewSqlStore(appCtx.GetDatabase())
	restaurantRatingStore := restaurantratingstore.NewSQLStore(appCtx.GetDatabase())

	go func() {
		defer common.AppRecovery()
		for {
			msg := <-c

			ratingData := msg.Data().(RatingData)

			rating, err := restaurantRatingStore.CalculateAVGPoint(ctx, map[string]interface{}{"restaurant_id": ratingData.GetRestaurantId()})
			if err != nil {
				log.Fatalln(err)
			}

			err = restaurantStore.UpdateRestaurantRating(ctx, ratingData.GetRestaurantId(), rating)
			if err != nil {
				log.Fatalln(err)
			}
		}
	}()
}
