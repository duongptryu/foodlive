package statisticmodel

type StatsUser struct {
	UserCount           []int    `json:"user_count"`
	Cate                []string `json:"cate"`
	NewMaxUser          int      `json:"new_max_user"`
	NewMinUser          int      `json:"new_min_user"`
	Growth              float64  `json:"growth"`
	GrowthWithLastMonth float64  `json:"growth_with_last_month"`
}
