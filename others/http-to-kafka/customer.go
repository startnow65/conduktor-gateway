package main

type Customer struct {
	Name          string  `json:"name"`
	Email         string  `json:"email"`
	FavouriteFood string  `json:"favourite_food"`
	Height        float32 `json:"height"`
}
