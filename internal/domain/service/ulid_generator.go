package service

import (
	"crypto/rand"
	"time"

	"github.com/oklog/ulid/v2"
)

type ULIDGenerator interface {
	Generate() string
}

type ulidGenerator struct {
}

func NewULIDGenerator() ULIDGenerator {
	return &ulidGenerator{}
}

func (u *ulidGenerator) Generate() string {
	return ulid.MustNew(ulid.Timestamp(time.Now().UTC()), rand.Reader).String()
}
