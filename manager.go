package session

import (
	"sync"
)

type Manager interface {
	Init(sessionId string, data map[string]interface{}) (*Session, error)
	Get(sessionId string) (*Session, error)
	Exist(sessionId string) (bool, error)
	Set(sessionId string, data map[string]interface{}) (*Session, error)
	Update(sessionId string, key string, value interface{}) (*Session, error)
	Delete(sessionId string) error
}

type manager struct {
	lock    sync.Mutex
	storage Storage
}

func NewManager(storage Storage) Manager {
	return &manager{
		storage: storage,
	}
}

func (m *manager) Exist(sessionId string) (bool, error) {
	return m.storage.Exist(sessionId)
}

func (m *manager) Init(sessionId string, data map[string]interface{}) (*Session, error) {
	m.lock.Lock()
	defer m.lock.Unlock()

	s := NewSession(sessionId)

	s.SetData(data)

	if err := m.storage.Set(sessionId, s); err != nil {
		return s, err
	}

	return m.storage.Get(sessionId)
}

func (m *manager) Get(sessionId string) (*Session, error) {
	return m.storage.Get(sessionId)
}

func (m *manager) Set(sessionId string, data map[string]interface{}) (*Session, error) {
	m.lock.Lock()
	defer m.lock.Unlock()

	s := NewSession(sessionId)
	s.SetData(data)

	if err := m.storage.Set(sessionId, s); err != nil {
		return s, err
	}

	return m.storage.Get(sessionId)
}

func (m *manager) Update(sessionId string, key string, value interface{}) (*Session, error) {
	m.lock.Lock()
	defer m.lock.Unlock()

	s, err := m.storage.Get(sessionId)
	if err != nil {
		return s, err
	}

	s.UpdateData(key, value)

	err = m.storage.Set(sessionId, s)
	if err != nil {
		return s, err
	}

	return m.storage.Get(sessionId)
}

func (m *manager) Delete(sessionId string) error {
	return m.storage.Delete(sessionId)
}
