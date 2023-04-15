package models

import "time"

type HouseInfo struct {
	Title string     `json:"title"`
	Link  string     `json:"link"`
	Date  *time.Time `json:"date"`
}
