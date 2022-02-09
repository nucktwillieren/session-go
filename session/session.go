package session

import "time"

type Session struct {
	Id         string                 `json:"id"`
	AccessedAt time.Time              `json:"accessedAt"`
	Data       map[string]interface{} `json:"data"`
}

func NewSession(id string) *Session {
	return &Session{
		Id: id,
	}
}

func (s *Session) SetData(data map[string]interface{}) {
	s.Data = data
}

func (s *Session) UpdateData(key string, value interface{}) {
	s.Data[key] = value
	s.AccessedAt = time.Now().UTC()
}

func (s *Session) GetData() map[string]interface{} {
	return s.Data
}

func (s *Session) GetId() string {
	return s.Id
}

func (s *Session) Get() Session {
	return Session{s.Id, s.AccessedAt, s.Data}
}
