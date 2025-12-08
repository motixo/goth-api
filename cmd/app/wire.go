//go:build wireinject

package main

//go:generate go run github.com/google/wire/cmd/wire@latest

import (
	"github.com/google/wire"
	"github.com/motixo/goat-api/internal/delivery/http"
	"github.com/motixo/goat-api/internal/di"
)

func InitializeApp() (*http.Server, error) {
	panic(wire.Build(di.ProviderSet))
}
