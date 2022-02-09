package storage

import (
	"encoding/base64"
	"encoding/json"

	"github.com/nucktwillieren/session-go/session"
)

func MarshalData(s *session.Session) (string, error) {
	data, err := json.Marshal(s)
	if err != nil {
		return "", err
	}

	sEnc := base64.StdEncoding.EncodeToString(data)

	return sEnc, nil
}

func UnmarshalData(s string) (*session.Session, error) {
	var entity session.Session

	sDec, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return &session.Session{}, err
	}

	err = json.Unmarshal(sDec, &entity)
	if err != nil {
		return &entity, err
	}

	return &entity, nil
}
