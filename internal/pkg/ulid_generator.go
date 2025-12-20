package pkg

import (
	"crypto/rand"
	"time"

	"github.com/oklog/ulid/v2"
)

func ULIDGenerator() string {
	return ulid.MustNew(ulid.Timestamp(time.Now().UTC()), rand.Reader).String()
}
