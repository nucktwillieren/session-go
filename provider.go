package session

type Provider interface {
	SetCookie(key interface{}, value interface{})
}
