package storage

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/nucktwillieren/auth/pkg/session"
)

type redisStorage struct {
	client      *redis.Client
	maxLifeTime int64
}

func NewRedisStorage(client *redis.Client, maxLifeTime int64) session.Storage {
	return &redisStorage{
		client:      client,
		maxLifeTime: maxLifeTime,
	}
}

func (rs *redisStorage) Exist(sessionId string) (bool, error) {
	_, err := rs.client.Get(context.Background(), sessionId).Result()
	if err == redis.Nil {
		return false, nil
	}

	if err != nil {
		return true, err
	}

	return true, nil
}

func (rs *redisStorage) Set(sessionId string, sessionEntity *session.Session) error {
	sData, err := MarshalData(sessionEntity)
	if err != nil {
		return err
	}

	_, err = rs.client.Set(context.Background(), sessionId, sData, time.Duration(rs.maxLifeTime)*time.Second).Result()
	if err != nil {
		return err
	}

	return nil
}

func (rs *redisStorage) Get(sessionId string) (*session.Session, error) {
	sData, err := rs.client.Get(context.Background(), sessionId).Result()
	if err != nil {
		return &session.Session{}, err
	}

	s, err := UnmarshalData(sData)
	if err != nil {
		return &session.Session{}, err
	}

	return s, nil
}

func (rs *redisStorage) GetAll() ([]session.Session, error) {
	return []session.Session{}, nil
}

func (rs *redisStorage) Delete(sessionId string) error {
	_, err := rs.client.Del(context.Background(), sessionId).Result()

	return err
}

func (rs *redisStorage) DeleteAll() error {
	return nil
}

func (rs *redisStorage) GC() error {
	return nil
}
