package subscribe

import (
	"context"
	"foodlive/common"
	"foodlive/component"
	"foodlive/modules/food/foodstore"
	"foodlive/modules/foodrating/foodratingstore"

	log "github.com/sirupsen/logrus"
)

func CalculateRatingFood(ctx context.Context, appCtx component.AppContext) {
	c, _ := appCtx.GetPubSubProvider().Subscribe(ctx, common.TopicUserCreateRestaurantRating)

	foodstore := foodstore.NewSqlStore(appCtx.GetDatabase())
	foodRatingStore := foodratingstore.NewSQLStore(appCtx.GetDatabase())

	go func() {
		defer common.AppRecovery()
		for {
			msg := <-c

			ratingData := msg.Data().(CastingFood)

			rating, err := foodRatingStore.CalculateAVGPoint(ctx, map[string]interface{}{"food_id": ratingData.GetFoodId()})
			if err != nil {
				log.Fatalln(err)
			}

			err = foodstore.UpdateRating(ctx, ratingData.GetFoodId(), rating)
			if err != nil {
				log.Fatalln(err)
			}
		}
	}()
}
