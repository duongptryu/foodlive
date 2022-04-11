package subscribe

import (
	"context"
	"foodlive/common"
	"foodlive/component"
	"foodlive/modules/food/foodstore"

	log "github.com/sirupsen/logrus"
)

type CastingFood interface {
	GetFoodId() int
}

func UserLikeFood(appCtx component.AppContext, ctx context.Context) {
	c, _ := appCtx.GetPubSubProvider().Subscribe(ctx, common.TopicUserLikeFood)

	foodStore := foodstore.NewSqlStore(appCtx.GetDatabase())

	go func() {
		defer common.AppRecovery()
		for {
			msg := <-c

			likeData := msg.Data().(CastingFood)

			err := foodStore.IncreaseLikeCount(ctx, likeData.GetFoodId())
			if err != nil {
				log.Fatalln(err)
			}
		}
	}()
}
