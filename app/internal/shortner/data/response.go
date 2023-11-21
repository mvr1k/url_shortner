package data

import "time"

type ShortUrl struct {
	Url      string `json:"url"`
	ExpireAt string `json:"expire_at"`
}

func NewShortUrl(url string) ShortUrl {
	return ShortUrl{
		url,
		time.Now().Add(1 * time.Minute).Format(time.DateTime),
	}
}
