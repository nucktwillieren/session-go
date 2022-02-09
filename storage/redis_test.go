package storage_test

import (
	"context"
	"log"
	"net"
	"net/url"
	"os"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/nucktwillieren/session-go/session"
	"github.com/nucktwillieren/session-go/storage"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/stretchr/testify/assert"
)

var (
	port           = "6379"
	client         *redis.Client
	sessionStorage session.Storage
)

func TestMain(m *testing.M) {
	wrapper := func(mm *testing.M) int {
		pool, err := dockertest.NewPool("")
		if err != nil {
			log.Fatalf("Could not connect to docker: %s", err)
		}
		opts := dockertest.RunOptions{
			Repository:   "redis",
			Env:          []string{},
			ExposedPorts: []string{"6379"},
			PortBindings: map[docker.Port][]docker.PortBinding{
				"6379": {
					{HostIP: "0.0.0.0", HostPort: port},
				},
			},
		}

		resource, err := pool.RunWithOptions(&opts)
		defer func(resource *dockertest.Resource) {
			if err := pool.Purge(resource); err != nil {
				log.Fatalf("Could not purge resource: %s", err)
			}
		}(resource)

		if err != nil {
			log.Fatalf("Could not start resource: %s", err.Error())
		}

		u, err := url.Parse(pool.Client.Endpoint())
		if err != nil {
			log.Fatalf("Could not parse endpoint: %s", pool.Client.Endpoint())
		}

		if err = pool.Retry(func() error {
			client = redis.NewClient(&redis.Options{
				Addr: net.JoinHostPort(u.Hostname(), resource.GetPort("6379/tcp")),
			})

			return client.Ping(context.Background()).Err()
		}); err != nil {
			log.Fatalf("Could not connect to redis: %s", err.Error())
		}

		if err = pool.Retry(func() error {
			sessionStorage = storage.NewRedisStorage(client, 600)

			return nil
		}); err != nil {
			log.Fatalf("Could not connect to docker: %s", err.Error())
		}

		return m.Run()
	}

	os.Exit(wrapper(m))
}

var (
	newSession = session.Session{
		Id:         "test",
		AccessedAt: time.Time{}.UTC(),
		Data: map[string]interface{}{
			"test2": "test2",
		},
	}
)

func TestRedisCreate(t *testing.T) {
	err := sessionStorage.Set(targetSession.Id, &targetSession)
	assert.NoError(t, err)
}

func TestRedisExist(t *testing.T) {
	ok, err := sessionStorage.Exist(targetSession.Id)
	assert.NotEmpty(t, ok)
	assert.NoError(t, err)
}

func TestRedisGet(t *testing.T) {
	s, err := sessionStorage.Get(targetSession.Id)
	assert.NotEmpty(t, s)
	assert.NoError(t, err)

	assert.EqualValues(t, targetSession, *s)
}

func TestRedisSet(t *testing.T) {
	err := sessionStorage.Set(targetSession.Id, &newSession)
	assert.NoError(t, err)

	s, err := sessionStorage.Get(targetSession.Id)
	assert.NotEmpty(t, s)
	assert.NoError(t, err)

	assert.EqualValues(t, newSession, *s)
}

func TestRedisDelete(t *testing.T) {
	err := sessionStorage.Delete(targetSession.Id)
	assert.NoError(t, err)

	ok, err := sessionStorage.Exist(targetSession.Id)
	assert.Empty(t, ok)
	assert.NoError(t, err)
}
