package main

type Promotion struct {
    Id string `json:"id"`
    Price float64 `json:"price"`
    Exp_date string `json:"expiration_date"`// todo convert date
}
