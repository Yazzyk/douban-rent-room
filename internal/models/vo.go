package models

import "time"

type HouseInfo struct {
	Title         string     `json:"title"`
	Link          string     `json:"link"`
	Author        string     `json:"author"`
	AuthorID      string     `json:"authorID"`
	AuthorLink    string     `json:"authorLink"`
	DataFrom      string     `json:"dataFrom"`
	Date          *time.Time `json:"date"`
	DateStr       string     `json:"dateStr"`
	CommentCount  int        `json:"commentCount"`
	AIReason      string     `json:"aiReason"`      // ai 推荐理由
	CommutingTime string     `json:"commutingTime"` // ai 通勤时间
}
