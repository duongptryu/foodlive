package subscribe

import (
	"context"
	"foodlive/common"
	"foodlive/component"
	"foodlive/modules/food/foodstore"

	log "github.com/sirupsen/logrus"
)

func DecreaseLikeCountFood(appCtx component.AppContext, ctx context.Context) {
	c, _ := appCtx.GetPubSubProvider().Subscribe(ctx, common.TopicUserUnlikeFood)

	foodStore := foodstore.NewSqlStore(appCtx.GetDatabase())

	go func() {
		defer common.AppRecovery()
		for {
			msg := <-c

			likeData := msg.Data().(CastingFood)

			err := foodStore.DecreaseLikeCount(ctx, likeData.GetFoodId())
			if err != nil {
				log.Fatalln(err)
			}
		}
	}()
}
