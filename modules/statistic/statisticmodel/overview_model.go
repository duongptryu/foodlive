package statisticmodel

type OverviewResp struct {
	TotalUser       int     `json:"total_user"`
	TodayUser       int     `json:"today_user"`
	TotalOrder      int     `json:"total_order"`
	TodayOrder      int     `json:"today_order"`
	TotalRestaurant int     `json:"total_restaurant"`
	TodayRestaurant int     `json:"today_restaurant"`
	TotalMoney      float64 `json:"total_money"`
}
