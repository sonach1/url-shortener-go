package models

import (
	"context"
	"urlshortener/config"
)

func SaveURL(short, original string) error {
	_, err := config.DB.Exec(`INSERT INTO urls (short, original) VALUES ($1, $2)`, short, original)
	return err
}

func GetOriginalURL(short string) (string, error) {
	var original string
	err := config.DB.QueryRow(`SELECT original FROM urls WHERE short = $1`, short).Scan(&original)
	return original, err
}

func CacheURL(short, original string) {
	config.RDB.Set(config.Ctx, short, original, 0)
}

func GetCachedURL(short string) (string, error) {
	return config.RDB.Get(config.Ctx, short).Result()
}
