package models

type GlobalScore struct {
	Title   string  `json:"title"`
	TitleTh string  `json:"title_th"`
	Average float64 `json:"average"`
	HiScore float64 `json:"hi_score"`
}
