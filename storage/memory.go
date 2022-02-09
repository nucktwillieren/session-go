package storage

import (
	"errors"

	"github.com/nucktwillieren/auth/pkg/session"
)

type memoryStorage struct {
	sessionList map[string]*session.Session
}

func NewMemoryStorage() session.Storage {
	return &memoryStorage{
		sessionList: make(map[string]*session.Session),
	}
}

func (ms *memoryStorage) Exist(sessionId string) (bool, error) {
	if _, ok := ms.sessionList[sessionId]; ok {
		return true, nil
	}

	return false, nil
}

func (ms *memoryStorage) Set(sessionId string, sessionEntity *session.Session) error {
	ms.sessionList[sessionId] = sessionEntity

	return nil
}

func (ms *memoryStorage) Get(sessionId string) (*session.Session, error) {
	if element, ok := ms.sessionList[sessionId]; ok {
		return element, nil
	}

	return &session.Session{}, errors.New("no such element")
}

func (ms *memoryStorage) GetAll() ([]session.Session, error) {
	return []session.Session{}, nil
}

func (ms *memoryStorage) Delete(sessionId string) error {
	if ok, _ := ms.Exist(sessionId); ok {
		delete(ms.sessionList, sessionId)
		return nil
	}

	return errors.New("no such element")
}

func (ms *memoryStorage) DeleteAll() error {
	return nil
}

func (ms *memoryStorage) GC() error {
	return nil
}
