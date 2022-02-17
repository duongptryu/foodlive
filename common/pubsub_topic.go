package common

import "foodlive/pubsub"

const (
	TopicUserLikeRestaurant         pubsub.Topic = "TopicUserLikeRestaurant"
	TopicUserDisLikeRestaurant      pubsub.Topic = "TopicUserDisLikeRestaurant"
	TopicUserCreateRestaurantRating pubsub.Topic = "TopicUserCreateRestaurantRating"
	TopicUserUpdateRestaurantRating pubsub.Topic = "TopicUserUpdateRestaurantRating"
)
