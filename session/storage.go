package session

type Storage interface {
	Set(sessionId string, sessionEntity *Session) error
	Exist(sessionId string) (bool, error)
	Get(sessionId string) (*Session, error)
	GetAll() ([]Session, error)
	Delete(sessionId string) error
	DeleteAll() error
	GC() error
}
