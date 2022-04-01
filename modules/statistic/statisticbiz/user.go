package statisticbiz

import (
	"context"
	"foodlive/modules/statistic/statisticmodel"
	"foodlive/modules/user/usermodel"
	"foodlive/modules/user/userstorage"
	"time"
)

type statsUserBiz struct {
	userStore userstorage.UserStore
}

func NewStatsUserBiz(userStore userstorage.UserStore) *statsUserBiz {
	return &statsUserBiz{
		userStore: userStore,
	}
}

func (biz *statsUserBiz) StatsUserBiz(ctx context.Context, year int) (*statisticmodel.StatsUser, error) {
	now := time.Now()
	yearCondition := time.Date(year, time.January, 1, 0, 0, 0, 0, now.Location())

	filter := usermodel.Filter{
		CreatedAtGt: &yearCondition,
	}
	result, err := biz.userStore.ListUserWithoutPaging(ctx, nil, &filter)
	if err != nil {
		return nil, err
	}

	var userCount []int
	var cate []string
	var check = make(map[string]int)

	for _, v := range result {
		//var key = v.CreatedAt.Format("2006-01-02")
		var key = v.CreatedAt.Month().String()
		if _, exist := check[key]; exist {
			check[key] = check[key] + 1
		} else {
			check[key] = 1
		}
	}

	var totaluser = 0
	var min = 100000
	var max = 0

	for d := yearCondition; d.After(time.Now()) == false; d = d.AddDate(0, 1, 0) {
		//k := d.Format("2006-01-02")
		k := d.Month().String()
		if v, exist := check[k]; exist {
			totaluser += v
			if min != 0 {
				min = v
			}
			if v > max {
				max = v
			}
			userCount = append(userCount, v)
		} else {
			min = 0
			userCount = append(userCount, 0)
		}
		cate = append(cate, k)
	}

	var growthWithLastMonth float64
	var growth float64
	if userCount[len(userCount)-2] == 0 {
		growth = float64((userCount[len(userCount)-1]) * 100)
	} else {
		growth = float64((userCount[len(userCount)-1] / userCount[len(userCount)-2]) * 100)
	}

	if userCount[0] == 0 {
		growth = float64((userCount[len(userCount)-1]) * 100)
	} else {
		growth = float64((userCount[len(userCount)-1] / userCount[0]) * 100)
	}

	return &statisticmodel.StatsUser{
		userCount,
		cate,
		max,
		min,
		growth,
		growthWithLastMonth,
	}, nil
}
