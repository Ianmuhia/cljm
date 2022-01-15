package services

import (
	"bytes"
	"context"
	"encoding/gob"
	"log"
	"time"

	redis_db "maranatha_web/datasources/redis"
	"maranatha_web/models"
)

var (
	CacheService cacheServiceInterface = &cacheService{}
)

type cacheService struct{}

type cacheServiceInterface interface {
	GetNewsList(ctx context.Context, id string) ([]models.News, error)
	SetNewsList(ctx context.Context, n []models.News) error
	SetNews(ctx context.Context, n *models.News) error
	GetNews(ctx context.Context, id string) (models.News, error)
}

func (c *cacheService) GetNewsList(ctx context.Context, id string) ([]models.News, error) {

	var news []models.News
	cmd := redis_db.RedisClient.Get(ctx, id)
	cmdb, err := cmd.Bytes()
	if err != nil {
		return news, err
	}
	b := bytes.NewReader(cmdb)
	if err := gob.NewDecoder(b).Decode(&news); err != nil {
		return news, err
	}
	return news, nil
}

func (c *cacheService) SetNewsList(ctx context.Context, n []models.News) error {
	var b bytes.Buffer
	if err := gob.NewEncoder(&b).Encode(n); err != nil {
		return err
	}
	return redis_db.RedisClient.Set(ctx, "news-list", b.Bytes(), 30*time.Second).Err()
}

func (c *cacheService) SetNews(ctx context.Context, n *models.News) error {
	var b bytes.Buffer
	if err := gob.NewEncoder(&b).Encode(n); err != nil {
		log.Println(err)
		return err
	}
	return redis_db.RedisClient.Set(ctx, "single-news", b.Bytes(), 30*time.Second).Err()
}

func (c *cacheService) GetNews(ctx context.Context, id string) (models.News, error) {
	var news models.News
	cmd := redis_db.RedisClient.Get(ctx, id)
	cmdb, err := cmd.Bytes()
	if err != nil {
		return news, err
	}
	b := bytes.NewReader(cmdb)
	if err := gob.NewDecoder(b).Decode(&news); err != nil {
		return news, err
	}
	return news, nil
}
