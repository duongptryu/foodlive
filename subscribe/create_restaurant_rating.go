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
	GetPoint() float64
}

func CalculateRatingRestaurantWhenUserCreate(ctx context.Context, appCtx component.AppContext) {
	c, _ := appCtx.GetPubSubProvider().Subscribe(ctx, common.TopicUserCreateRestaurantRating)

	restaurantStore := restaurantstore.NewSqlStore(appCtx.GetDatabase())
	restaurantRatingStore := restaurantratingstore.NewSQLStore(appCtx.GetDatabase())

	go func() {
		defer common.AppRecovery()

		msg := <-c

		ratingData := msg.Data().(RatingData)

		countRated, err := restaurantRatingStore.CountRestaurantRating(ctx, map[string]interface{}{"restaurant_id": ratingData.GetRestaurantId()})
		if err != nil {
			log.Error(err)
		}

		rstDb, err := restaurantStore.FindRestaurant(ctx, map[string]interface{}{"id": ratingData.GetRestaurantId()})
		if err != nil {
			log.Error(err)
		}

		newRateNumber := (rstDb.Rating + ratingData.GetPoint()) / float64(countRated)

		err = restaurantStore.UpdateRestaurantRating(ctx, ratingData.GetRestaurantId(), newRateNumber)
		if err != nil {
			log.Error(err)
		}
	}()
}
