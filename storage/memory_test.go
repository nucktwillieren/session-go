package storage_test

import (
	"testing"

	"github.com/nucktwillieren/session-go/storage"
	"github.com/stretchr/testify/assert"
)

var (
	memory = storage.NewMemoryStorage()
)

func TestMemoryCreate(t *testing.T) {
	err := memory.Set(targetSession.Id, &targetSession)
	assert.NoError(t, err)
}

func TestMemoryExist(t *testing.T) {
	ok, err := memory.Exist(targetSession.Id)
	assert.NotEmpty(t, ok)
	assert.NoError(t, err)
}

func TestMemoryGet(t *testing.T) {
	s, err := memory.Get(targetSession.Id)
	assert.NotEmpty(t, s)
	assert.NoError(t, err)

	assert.EqualValues(t, targetSession, *s)
}

func TestMemorySet(t *testing.T) {
	err := memory.Set(targetSession.Id, &newSession)
	assert.NoError(t, err)

	s, err := memory.Get(targetSession.Id)
	assert.NotEmpty(t, s)
	assert.NoError(t, err)

	assert.EqualValues(t, newSession, *s)
}

func TestMemoryDelete(t *testing.T) {
	err := memory.Delete(targetSession.Id)
	assert.NoError(t, err)

	ok, err := memory.Exist(targetSession.Id)
	assert.Empty(t, ok)
	assert.NoError(t, err)
}
