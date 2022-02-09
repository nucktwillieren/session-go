package storage_test

import (
	"testing"
	"time"

	"github.com/nucktwillieren/auth/pkg/session"
	"github.com/nucktwillieren/auth/pkg/session/storage"
	"github.com/stretchr/testify/assert"
)

var (
	targetSession = session.Session{
		Id:         "test",
		AccessedAt: time.Time{}.UTC(),
		Data: map[string]interface{}{
			"test": "test",
		},
	}
	jsonBase64String = "eyJpZCI6InRlc3QiLCJhY2Nlc3NlZEF0IjoiMDAwMS0wMS0wMVQwMDowMDowMFoiLCJkYXRhIjp7InRlc3QiOiJ0ZXN0In19"
)

func TestMarshalData(t *testing.T) {
	se, err := storage.MarshalData(&targetSession)
	assert.NotEmpty(t, se)
	assert.NoError(t, err)

	assert.Equal(t, jsonBase64String, se)
}

func TestUnmarshalData(t *testing.T) {
	s, err := storage.UnmarshalData(jsonBase64String)
	assert.NotEmpty(t, s)
	assert.NoError(t, err)

	assert.EqualValues(t, targetSession, *s)
}

func TestMarshalAndUnmarshal(t *testing.T) {
	se, err := storage.MarshalData(&targetSession)
	assert.NotEmpty(t, se)
	assert.NoError(t, err)

	s, err := storage.UnmarshalData(se)
	assert.NotEmpty(t, s)
	assert.NoError(t, err)

	assert.EqualValues(t, targetSession, *s)
}
