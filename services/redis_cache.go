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
	// GetNewsList for cache
	GetNewsList(ctx context.Context, id string) ([]models.News, error)
	SetNewsList(ctx context.Context, n []models.News) error
	SetNews(ctx context.Context, n *models.News) error
	GetNews(ctx context.Context, id string) (models.News, error)

	// GetPartnersList for cache
	GetPartnersList(ctx context.Context, id string) (interface{}, error)
	SetPartnersList(key string, ctx context.Context, n interface{}) error

	// GetBooksList  for cache
	GetBooksList(ctx context.Context, id string) ([]models.Books, error)
	GetBooks(ctx context.Context, id string) (models.Books, error)
	SetBooks(ctx context.Context, n *models.Books) error

	//GetPrayerList for cache
	GetPrayerList(ctx context.Context, id string) ([]models.Prayer, error)
	GetPrayer(ctx context.Context, id string) (models.Prayer, error)
	SetPrayer(ctx context.Context, n *models.Prayer) error

	//GetTestimoniesList for cache
	GetTestimoniesList(ctx context.Context, id string) ([]models.Testimonies, error)
	GetTestimonies(ctx context.Context, id string) (models.Testimonies, error)
	SetTestimonies(ctx context.Context, n *models.Testimonies) error

	//GetEventsList for cache
	GetEventsList(ctx context.Context, id string) ([]models.Events, error)
	GetEvents(ctx context.Context, id string) (models.Events, error)
	SetEvents(ctx context.Context, n *models.Events) error

	//GetGenreList
	GetGenreList(ctx context.Context, id string) ([]models.Genre, error)
	GetGenre(ctx context.Context, id string) (models.Genre, error)
	SetGenre(ctx context.Context, n *models.Genre) error
}

// News/***
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

// Partners /***
func (c *cacheService) GetPartner(key string, ctx context.Context, n interface{}) error {
	var b bytes.Buffer
	if err := gob.NewEncoder(&b).Encode(n); err != nil {
		log.Println(err)
		return err
	}
	return redis_db.RedisClient.Set(ctx, key, b.Bytes(), 1*time.Minute).Err()
}
func (c *cacheService) SetPartnersList(key string, ctx context.Context, n interface{}) error {
	var b bytes.Buffer
	if err := gob.NewEncoder(&b).Encode(n); err != nil {
		return err
	}
	return redis_db.RedisClient.Set(ctx, key, b.Bytes(), 1*time.Minute).Err()
}
func (c *cacheService) GetPartnersList(ctx context.Context, id string) (interface{}, error) {
	var data interface{}

	cmd := redis_db.RedisClient.Get(ctx, id)

	cmdb, err := cmd.Bytes()
	if err != nil {
		return data, err
	}
	b := bytes.NewReader(cmdb)
	if err := gob.NewDecoder(b).Decode(&data); err != nil {
		log.Println(err)
		return data, err
	}

	return data, nil
}

// Books /***
func (c *cacheService) GetBooksList(ctx context.Context, id string) ([]models.Books, error) {
	var books []models.Books
	cmd := redis_db.RedisClient.Get(ctx, id)
	cmdb, err := cmd.Bytes()
	if err != nil {
		return books, err
	}
	b := bytes.NewReader(cmdb)
	if err := gob.NewDecoder(b).Decode(&books); err != nil {
		return books, err
	}
	return books, nil
}
func (c *cacheService) GetBooks(ctx context.Context, id string) (models.Books, error) {
	var books models.Books
	cmd := redis_db.RedisClient.Get(ctx, id)
	cmdb, err := cmd.Bytes()
	if err != nil {
		return books, err
	}
	b := bytes.NewReader(cmdb)
	if err := gob.NewDecoder(b).Decode(&books); err != nil {
		return books, err
	}
	return books, nil
}
func (c *cacheService) SetBooks(ctx context.Context, n *models.Books) error {
	var b bytes.Buffer
	if err := gob.NewEncoder(&b).Encode(n); err != nil {
		log.Println(err)
		return err
	}
	return redis_db.RedisClient.Set(ctx, "single-books", b.Bytes(), 30*time.Second).Err()
}

// Prayer /***
func (c *cacheService) GetPrayerList(ctx context.Context, id string) ([]models.Prayer, error) {
	var prayer []models.Prayer
	cmd := redis_db.RedisClient.Get(ctx, id)
	cmdb, err := cmd.Bytes()
	if err != nil {
		return prayer, err
	}
	b := bytes.NewReader(cmdb)
	if err := gob.NewDecoder(b).Decode(&prayer); err != nil {
		return prayer, err
	}
	return prayer, nil
}
func (c *cacheService) GetPrayer(ctx context.Context, id string) (models.Prayer, error) {
	var prayer models.Prayer
	cmd := redis_db.RedisClient.Get(ctx, id)
	cmdb, err := cmd.Bytes()
	if err != nil {
		return prayer, err
	}
	b := bytes.NewReader(cmdb)
	if err := gob.NewDecoder(b).Decode(&prayer); err != nil {
		return prayer, err
	}
	return prayer, nil
}
func (c *cacheService) SetPrayer(ctx context.Context, n *models.Prayer) error {
	var b bytes.Buffer
	if err := gob.NewEncoder(&b).Encode(n); err != nil {
		log.Println(err)
		return err
	}
	return redis_db.RedisClient.Set(ctx, "single-prayer", b.Bytes(), 30*time.Second).Err()
}

// Testimonies /***
func (c *cacheService) GetTestimoniesList(ctx context.Context, id string) ([]models.Testimonies, error) {
	var testimonies []models.Testimonies
	cmd := redis_db.RedisClient.Get(ctx, id)
	cmdb, err := cmd.Bytes()
	if err != nil {
		return testimonies, err
	}
	b := bytes.NewReader(cmdb)
	if err := gob.NewDecoder(b).Decode(&testimonies); err != nil {
		return testimonies, err
	}
	return testimonies, nil
}
func (c *cacheService) GetTestimonies(ctx context.Context, id string) (models.Testimonies, error) {
	var testimonies models.Testimonies
	cmd := redis_db.RedisClient.Get(ctx, id)
	cmdb, err := cmd.Bytes()
	if err != nil {
		return testimonies, err
	}
	b := bytes.NewReader(cmdb)
	if err := gob.NewDecoder(b).Decode(&testimonies); err != nil {
		return testimonies, err
	}
	return testimonies, nil
}
func (c *cacheService) SetTestimonies(ctx context.Context, n *models.Testimonies) error {
	var b bytes.Buffer
	if err := gob.NewEncoder(&b).Encode(n); err != nil {
		log.Println(err)
		return err
	}
	return redis_db.RedisClient.Set(ctx, "single-testimonies", b.Bytes(), 30*time.Second).Err()
}

// Events /***
func (c *cacheService) GetEventsList(ctx context.Context, id string) ([]models.Events, error) {
	var events []models.Events
	cmd := redis_db.RedisClient.Get(ctx, id)
	cmdb, err := cmd.Bytes()
	if err != nil {
		return events, err
	}
	b := bytes.NewReader(cmdb)
	if err := gob.NewDecoder(b).Decode(&events); err != nil {
		return events, err
	}
	return events, nil
}
func (c *cacheService) GetEvents(ctx context.Context, id string) (models.Events, error) {
	var events models.Events
	cmd := redis_db.RedisClient.Get(ctx, id)
	cmdb, err := cmd.Bytes()
	if err != nil {
		return events, err
	}
	b := bytes.NewReader(cmdb)
	if err := gob.NewDecoder(b).Decode(&events); err != nil {
		return events, err
	}
	return events, nil
}
func (c *cacheService) SetEvents(ctx context.Context, n *models.Events) error {
	var b bytes.Buffer
	if err := gob.NewEncoder(&b).Encode(n); err != nil {
		log.Println(err)
		return err
	}
	return redis_db.RedisClient.Set(ctx, "single-events", b.Bytes(), 30*time.Second).Err()
}

// Genre /***
func (c *cacheService) GetGenreList(ctx context.Context, id string) ([]models.Genre, error) {
	var genre []models.Genre
	cmd := redis_db.RedisClient.Get(ctx, id)
	cmdb, err := cmd.Bytes()
	if err != nil {
		return genre, err
	}
	b := bytes.NewReader(cmdb)
	if err := gob.NewDecoder(b).Decode(&genre); err != nil {
		return genre, err
	}
	return genre, nil
}
func (c *cacheService) GetGenre(ctx context.Context, id string) (models.Genre, error) {
	var genre models.Genre
	cmd := redis_db.RedisClient.Get(ctx, id)
	cmdb, err := cmd.Bytes()
	if err != nil {
		return genre, err
	}
	b := bytes.NewReader(cmdb)
	if err := gob.NewDecoder(b).Decode(&genre); err != nil {
		return genre, err
	}
	return genre, nil
}
func (c *cacheService) SetGenre(ctx context.Context, n *models.Genre) error {
	var b bytes.Buffer
	if err := gob.NewEncoder(&b).Encode(n); err != nil {
		log.Println(err)
		return err
	}
	return redis_db.RedisClient.Set(ctx, "single-genre", b.Bytes(), 30*time.Second).Err()
}
