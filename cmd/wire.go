//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"qyx/internal/controllers"
	"qyx/internal/services"
	"qyx/pkg/start"
)

var providerSet = wire.NewSet(
	NewGinEngine,
	start.NewGinStart,
	controllers.ProviderSet,
	services.ProviderSet,
)

func CreateApp() *start.GinStart {
	panic(wire.Build(providerSet))
}
